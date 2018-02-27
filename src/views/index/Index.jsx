import React, { Component } from 'react'
import { NavLink } from 'react-router-dom'
import './style.scss'
import API_CONFIG from '../../api'
import Sidebar from '../../components/sidebar/Sidebar'
import { querystring } from '../../utils'
import { Pagination } from 'antd'
import List from '../../components/topics-list/TopicsList'
/* eslint-disable */
class HomePage extends Component {

    constructor (props) {
        super(props)
        this.state = {
            page: 1,        // 当前页
            total: 9999,    // 总条数
            topics: [],     // 主题列表
            mark: false,
        }
    }

    componentDidMount () {
        this.setState({
            page: parseInt(querystring(this.props.location.search).page) || 1,
        }, () => {
            this.fetchTopics();
        })
    }

    componentDidUpdate (prevProps, prevState, prevContext) {
        if( this.props.location != prevProps.location ) {
            var page = parseInt(querystring(this.props.location.search).page);
            if( !page ) {
                this.setState({page: 1}, () => {
                    this.fetchTopics();
                });
                return;
            }
            this.fetchTopics();
        }
    }

    /**
     * @func 获取主题列表
     * @param {Number} page
     * @param {String} tab
     */
    fetchTopics = () => {
        this.setState({
            mark: true,
        });
        var beforeTime = Date.now();
        axios.get(API_CONFIG.topics, {
            params: {
                limit: 40,
                mdrender: false,
                tab: querystring(this.props.location.search).tab || 'all',
                page: this.state.page,
            }
        })
        .then(res => {
            var afterTime = Date.now() - beforeTime;
            if( afterTime <= 300 ) {
                setTimeout(() => {
                    this.setState({
                        mark: false,
                    });   
                }, 300 - afterTime)
            } else {
                this.setState({
                    mark: false,
                });
            }
            if( res.data.success ) {
                this.setState({
                    topics: res.data.data
                });
            } 
        })
        .catch(e => e);
    }

    isActive (tabVal) {
        return querystring(this.props.location.search).tab === tabVal;
    }

    homePageActive = () => {
        var tab = querystring(this.props.location.search).tab;
        return !tab || tab == 'all';
    }

    currentChange = (page) => {
        this.setState({
            page,
        })
        var tab = querystring(this.props.location.search).tab || 'all';
        this.props.history.push({
            pathname: '/',
            search: `?tab=${tab}&page=${page}`,
        });
        window.scrollTo(0, 0);
    }

    render () {
        return (
            <section className="index-section">
                <div className="topics-container index-container">
                    {/* 导航 */}
                    <nav className="nav">
                        <NavLink to="/" isActive={this.homePageActive}>全部</NavLink>
                        <NavLink to="/?tab=good" isActive={this.isActive.bind(this, 'good')}>精华</NavLink>
                        <NavLink to="/?tab=share" isActive={this.isActive.bind(this, 'share')}>分享</NavLink>
                        <NavLink to="/?tab=ask" isActive={this.isActive.bind(this, 'ask')}>问答</NavLink>
                        <NavLink to="/?tab=job" isActive={this.isActive.bind(this, 'job')}>招聘</NavLink>
                        <NavLink to="/?tab=dev" isActive={this.isActive.bind(this, 'dev')}>客户端测试</NavLink>
                    </nav>
                    <div className="topics-list">
                        <div className="mark-box" style={{display: !this.state.mark ? 'none' : ''}}>
                            <div className="mark-line"></div>
                            <div className="mark-line"></div>
                            <div className="mark-line"></div>
                        </div>
                        <List topics={this.state.topics} />
                    </div>
                    <div className="pagination-box">
                        <Pagination 
                            current={this.state.page} 
                            onChange={this.currentChange} 
                            total={this.state.total} 
                            pageSize={40} />
                    </div>
                </div>
                {/* 侧边栏 */}
                <Sidebar />
            </section>
        );
    }
}

export default HomePage;