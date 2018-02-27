import React, { Component } from 'react'
import './collections.scss'
import Sidebar from '../../components/sidebar/Sidebar'
import { Link } from 'react-router-dom'
import API_CONFIG from '../../api'
import { message } from 'antd'
import List from '../../components/topics-list/TopicsList'
/* eslint-disable */
class Collections extends Component {

    constructor (props) {
        super(props);
        this.state = {
            userCollections: [],
            loading: true,
        }
    }

    fetchCollections = () => {
        axios.get(`${API_CONFIG.userCollections}${this.props.match.params.loginname}`)
        .then(res => {
            if( res.data.success ) {
                this.setState({
                    userCollections: res.data.data,
                    loading: false,
                })
            }
        })
        .catch(e => {
            message.warning('不存在此用户');
            this.props.history.replace('/');
        });
    }

    componentDidMount () {
        this.fetchCollections();
    }

    render () {
        return (
            <section className="index-section">
                <div className="topics-container collections">
                    {
                        this.state.loading &&
                        <div className="collections-loading">loading...</div>
                    }
                    <div className="collections-title">
                        <Link to="/">主页</Link>
                        <em className="slashes"> / </em>
                        <span>{ this.props.match.params.loginname } 收藏的话题</span>
                    </div>
                    <List topics={this.state.userCollections} />
                </div>
                <Sidebar />
            </section>
        )
    }
}

export default Collections;






