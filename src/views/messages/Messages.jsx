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
    @observable loading = true;

    constructor (props) {
        super(props);
        this.state = {};
    }

    // 获取已读和未读消息
    fetchMessages () {
        axios.get(API_CONFIG.fetchMessages)
        .then(res => {
            if( res.data.success ) {
                this.loading = false;
                this.message = res.data.data;
            }
        })
        .catch(e => {
            this.props.history.replace('/');
        });
    }

    componentDidMount () {
        this.fetchMessages();
        // 标记全部已读
        if( this.props.store.messageCount > 0 ) {
            axios.post(API_CONFIG.messageMarkAll)
            .then(res => {})
            .catch(e => e);
        }
    }

    render () {
        // 消息组件
        var Msg = ({ obj }) => {
            return (
                <div className="msg-list">
                    { this.loading && <div className="msg-loading"></div> }
                    { this.message[obj].length === 0 && <div className="no-msg">暂无消息</div> }
                    <ul>
                    {
                        this.message[obj].map((item, index) => {
                            return (
                                <li key={item.id}>
                                    {
                                        item.type === 'reply' ?
                                        <div>
                                            <Link to={`/user/${item.author.loginname}`}>{ item.author.loginname }</Link>
                                            <span> 回复了你的话题 </span>
                                            <Link to={`/topic/${item.topic.id}`}>{ item.topic.title }</Link>
                                        </div>
                                        : item.type == 'at'
                                        ?
                                        <div>
                                            <Link to={`/user/${item.author.loginname}`}>{ item.author.loginname }</Link>
                                            <span> 在话题 </span>
                                            <Link to={`/topic/${item.topic.id}`}>{ item.topic.title }</Link>
                                            <span> 中@了你</span>
                                        </div>
                                        : ''
                                    }
                                </li>
                            );
                        })
                    }
                    </ul>
                </div>
            );
        };
        return (
            <section className="index-section">
                <div className="topics-container messages">
                    <div className="new-msg">
                        <div className="top">
                            <Link to="/">主页</Link>
                            <em> / </em>
                            <span>新消息</span>
                        </div>
                        <Msg obj="hasnot_read_messages" />
                    </div>
                    <div className="past-times">
                        <div className="top">已读消息</div>
                        <Msg obj="has_read_messages" />
                    </div>
                </div>
                <Sidebar />
            </section>
        )
    }
}


export default Messages;

