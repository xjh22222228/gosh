import React, { Component } from 'react'
import Sidebar from '../../components/sidebar/Sidebar'
import API_CONFIG from '../../api'
import { message } from 'antd'
import { Link } from 'react-router-dom'
import './user.scss'
import { fromNow } from '../../utils'
import List from '../../components/topics-list/TopicsList'

/* eslint-disable */
class User extends Component {

    constructor (props) {
        super(props);
        this.state = {
            user: {          // 用户详情
                loginname: '--',
                avatar_url: 'data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAMAAAACAQMAAACnuvRZAAAAA1BMVEX29vYACyOqAAAACklEQVQI12MAAgAABAABINItbwAAAABJRU5ErkJggg==',
                githubUsername: 'alsotang',
                create_at: Date.now(),
                score: 0,
                recent_topics: [],              // 最近创建的话题
                recent_replies: [],             // 最近参与的话题
            }
        };
    }

    // 获取用户详情
    fetchUserDetail = () => {
        axios.get(`${API_CONFIG.user}${this.props.match.params.loginname}`)
        .then(res => {
            if( res.data.success ) {
                this.setState({
                    user: res.data.data
                });
            }
        })
        .catch(e => {
            message.warning('不存在此用户');
            this.props.history.replace('/');
        })
    }
    
    componentDidMount () {
        this.fetchUserDetail();
    }

    render () {
        var user = this.state.user;
        return (
            <section className="index-section">
                <div className="topics-container user">
                    <div className="box">
                        <div className="box-title">
                            <Link to="/">主页</Link>
                            <em className="slashes"> / </em>
                            <span>个人主页</span>
                        </div>
                        <div className="user-info">
                            <div className="user">
                                <img src={user.avatar_url} alt="avatar" />
                                <span>{ user.loginname }</span>
                            </div>
                            <div>{ user.score } 积分</div>
                            <div className="view-topics-collections">
                                <Link to={`/user/${user.loginname}/collections`}>查看话题收藏</Link>
                            </div>
                            <div className="create-at">注册时间 { fromNow(user.create_at) }</div>
                        </div>
                    </div>
                    <div className="box">
                        <div className="box-title">最近创建的话题</div>
                        <List topics={user.recent_topics} />
                    </div>
                    <div className="box">
                        <div className="box-title">最近参与的话题</div>
                        <List topics={user.recent_replies} />
                    </div>
                </div>
                <Sidebar />
            </section>
        )
    }
}

export default User;


