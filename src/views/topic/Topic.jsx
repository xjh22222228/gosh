import React, { Component } from 'react'
import Sidebar from '../../components/sidebar/Sidebar'
import API_CONFIG from '../../api'
import { fromNow } from '../../utils'
import { observer, inject } from 'mobx-react'
import './style.scss'
import { Icon, message } from 'antd'
import { observable } from 'mobx'

/* eslint-disable */
@inject(stores => stores)
@observer class Topic extends Component {

    @observable loading = false;
    @observable detail = {
        author: {
            avatar_url: 'https://avatars1.githubusercontent.com/u/15535177?s=460&v=4',
            loginname: '--',
        },
        author_id: '',
        content: '',
        create_at: Date.now(),
        good: false,
        id: '',
        is_collect: false,
        last_reply_at: Date.now(),
        replies: [],
        reply_count: 0,
        tab: '',
        title: '',
        top: false,
        visit_count: 0,
    }

    constructor (props) {
        super(props);
        this.state = {};
    }

    comeFrom = tab => {
        switch (tab) {
            case 'ask':
                return '问答';
            case 'share':
                return '分享';
            case 'job':
                return '招聘';
            case 'good':
                return '精华';
            default:
                return '';
        }
    }

    tag = item => {
        if( item.top ) return '置顶';
        if( item.good ) return '精华';
    }

    // 获取主题详情
    fetchTopic = () => {
        this.loading = true;
        axios.get(`${API_CONFIG.topicDetail}${this.props.match.params.id}`)
        .then(res => {
            this.loading = false;
            if( res.data.success ) this.detail = res.data.data;
        })
        .catch(e => {
            this.loading = false;
            this.props.history.replace('/');
        });
    }

    // 收藏 && 取消收藏主题
    collectionBtn = () => {
        var url = this.detail.is_collect ? API_CONFIG.cancelCollection : API_CONFIG.collection;
        axios.post(url, {
            topic_id: this.detail.id
        })
        .then(res => {
            if( res.data.success ) {
                this.detail.is_collect = !this.detail.is_collect;
            }
        })
        .catch(e => e);
    }

    // 点赞/取消点赞
    likeBtn (id, index) {
        if( !this.props.store.isLogin ) return message.warning('您未登陆!');;
        axios.post(`${API_CONFIG.like}${id}/ups`)
        .then(res => {
            if( res.data.success ) {
                // 取消赞
                if( res.data.action === 'down' ) {
                    this.detail.replies[index].is_uped = false;
                    this.detail.replies[index].ups.pop();
                } else {
                    this.detail.replies[index].is_uped = true;
                    this.detail.replies[index].ups.push(Date.now());
                }
            }
        })
        .catch(e => e);
    }

    componentDidMount () {
        this.fetchTopic();
        window.scrollTo(0, 0);
    }

    render () {
        return (
            <section className="topic index-section">
                <div className="topics-container">
                    <div className="detail">
                        {
                            this.loading && <div className="loading">loading...</div>
                        }
                        <div className="topic-top">
                            <div className="topic-title">
                                {
                                    (this.detail.top || this.detail.good) &&
                                    <span className="tag">{ this.tag(this.detail) }</span>
                                }
                                <h1>{ this.detail.title }</h1>
                            </div>
                            <div className="topic-bottom">
                                <div className="topic-info">
                                    <span>• 发布于 { fromNow(this.detail.create_at) } • 作者 </span>
                                    <a href={`https://cnodejs.org/user/${this.detail.author.loginname}`}>{ this.detail.author.loginname }</a>
                                    <span> • { this.detail.visit_count } 次浏览 • 最后一次编辑是 { fromNow(this.detail.last_reply_at) } • 来自 { this.comeFrom(this.detail.tab) }</span>
                                </div>
                                {
                                    this.props.store.isLogin
                                    &&
                                    <div className="collection user-select-none">
                                        <button onClick={this.collectionBtn}>{ this.detail.is_collect ? '取消收藏' : '收藏' }</button>
                                    </div>
                                }
                            </div>
                        </div>
                        <div className="content markdown-body" dangerouslySetInnerHTML={{__html: this.detail.content}}></div>
                    </div>
                    {/* 回复区域 */}
                    {
                        this.detail.reply_count > 0
                        &&
                        <div className="reply">
                            <div className="reply-count">{ this.detail.reply_count } 回复</div>
                            <ul>
                                {
                                    this.detail.replies.map((item, index) => {
                                        return (
                                            <li key={item.id}>
                                                <div className="avatar">
                                                    <img src={ item.author.avatar_url } alt="头像" />
                                                </div>
                                                <div className="reply-right">
                                                    <div className="reply-author">
                                                        <em>{ item.author.loginname }</em>
                                                        <span>{ index + 1 }楼•{ fromNow(item.create_at) }</span>
                                                        {
                                                           this.detail.author.loginname == item.author.loginname && <strong>作者</strong>
                                                        }
                                                    </div>
                                                    <div className="operation user-select-none">
                                                        <div>
                                                            <Icon type={item.is_uped ? 'like' : 'like-o'} onClick={this.likeBtn.bind(this, item.id, index)} />
                                                            <em>{ item.ups.length }</em>
                                                        </div>
                                                    </div>
                                                </div>
                                                <div className="reply-content markdown-body" dangerouslySetInnerHTML={{__html: item.content}}></div>
                                            </li>
                                        );
                                    })
                                }
                            </ul>
                        </div>
                    }
                </div>
                <Sidebar author={this.detail.author} from="topic" />
            </section>
        )
    }
}



export default Topic;






