import React, { Component } from 'react'
import { observer, inject } from 'mobx-react'
import { observable } from 'mobx'
import Sidebar from '../../components/sidebar/Sidebar'
import API_CONFIG from '../../api'
import { Link } from 'react-router-dom'
import './messages.scss'
/* eslint-disable */
@inject(stores => stores)
@observer class Messages extends Component {

    @observable message = {         // 已读和未读消息
        has_read_messages: [],
        hasnot_read_messages: []
    };

    constructor (props) {
        super(props);
        this.state = {};
    }

    // 获取已读和未读消息
    fetchMessages () {
        axios.get(API_CONFIG.fetchMessages)
        .then(res => {
            if( res.data.success ) {
                this.message = res.data.data;
            }
        })
        .catch(e => {
            this.props.history.replace('/');
        });
    }

    componentDidMount () {
        this.fetchMessages();
    }

    render () {
        return (
            <section className="index-section">
                <div className="topics-container messages">
                    <div className="new-msg">
                        <div className="top">
                            <Link to="/">主页</Link>
                            <em> / </em>
                            <span>新消息</span>
                        </div>
                        <div className="msg-list">
                            {
                                this.message.hasnot_read_messages.length === 0 && <div className="no-msg">暂无消息</div>
                            }
                            <ul>
                            {
                                this.message.hasnot_read_messages.map((item, index) => {
                                    return (
                                        <li key={item.id}>
                                            <a href={`https://cnodejs.org/user/${item.author.loginname}`}>{ item.author.loginname }</a>
                                            <span> 回复了你的话题 </span>
                                            <Link to={`/topic/${item.topic.id}`}>{ item.topic.title }</Link>
                                        </li>
                                    );
                                })
                            }
                            </ul>
                        </div>
                    </div>
                    <div className="past-times">
                        <div className="top">已读消息</div>
                        <div className="msg-list">
                            {
                                this.message.has_read_messages.length === 0 && <div className="no-msg">暂无消息</div>
                            }
                            <ul>
                            {
                                this.message.has_read_messages.map((item, index) => {
                                    return (
                                        <li key={item.id}>
                                            <a href={`https://cnodejs.org/user/${item.author.loginname}`}>{ item.author.loginname }</a>
                                            <span> 回复了你的话题 </span>
                                            <Link to={`/topic/${item.topic.id}`}>{ item.topic.title }</Link>
                                        </li>
                                    );
                                })
                            }
                            </ul>
                        </div>
                    </div>
                </div>
                <Sidebar />
            </section>
        )
    }
}


export default Messages;

