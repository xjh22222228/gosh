import React, { Component } from 'react'
import Sidebar from '../../components/sidebar/Sidebar'
import API_CONFIG from '../../api'
import { fromNow } from '../../utils'
import { observer, inject } from 'mobx-react'
import './topic.scss'
import { Icon, message } from 'antd'
import { observable } from 'mobx'
import SimpleMDE from 'simplemde'
import { Link } from 'react-router-dom'

/* eslint-disable */
@inject(stores => stores)
@observer class Topic extends Component {

    @observable loading = true;
    @observable detail = {
        author: {
            avatar_url: 'data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAMAAAACAQMAAACnuvRZAAAAA1BMVEX29vYACyOqAAAACklEQVQI12MAAgAABAABINItbwAAAABJRU5ErkJggg==',
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
    @observable insertBtnText = '回复';

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
        axios.get(`${API_CONFIG.topicDetail}${this.props.match.params.id}`)
        .then(res => {
            if( res.data.success ) {
                this.detail = res.data.data;
                this.loading = false;
            }
        })
        .catch(e => {
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

    /**
     * @func 点赞/取消点赞
     * @param {String} id 
     * @param {Object} author 
     * @param {Number} index 
     */
    likeBtn (id, author, index) {
        try {
            if( !this.props.store.isLogin ) throw new Error('您未登陆!');
            if( author.loginname === this.props.store.userInfo.loginname ) throw new Error('不能赞自己!');
        } catch (e) {
            return message.warning(e.message);
        }
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

    // 初始化markdown编辑器
    initMarkdownEditor = () => {
        this.simplemde = new SimpleMDE({
            element: document.getElementById("markdown-editor"),
            spellChecker: false, 				// 启用拼写检查，会有背景色
            autoDownloadFontAwesome: false,		// 是否需要下载字体图标
        });
    }

    // 插入评论
    insertReply = () => {
        // 避免多次点击按钮
        if( this.insertBtnText === '发送中...' ) return;
        var val = this.simplemde.value();
        if( !val ) return message.warning('内容不能为空!');
        this.insertBtnText = '发送中...';
        axios.post(`${API_CONFIG.replies}${this.detail.id}/replies`, {
            content: val,
        })
        .then(res => {
            if( res.data.success ) {
                this.simplemde.value('');
                message.success('发送成功!');
                this.insertBtnText = '回复';
                // 更新资源
                this.fetchTopic();
            }
        })
        .catch(e => {
            this.insertBtnText = '回复';
        });
    }

    // 回复别人
    replyOthers (loginname) {
        var top = document.querySelector('.insert-reply').offsetTop;
        window.scrollTo(0, top - 80);
        this.simplemde.value(`@${loginname} `);
    }

    componentDidMount () {
        this.fetchTopic();
        this.initMarkdownEditor();
        window.scrollTo(0, 0);
    }

    render () {
        return (
            <section className="topic index-section">
                <div className="topics-container">
                    <div className="detail">
                        { this.loading && <div className="loading">loading...</div> }
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
                                    <Link to={`/user/${this.detail.author.loginname}`}>{ this.detail.author.loginname }</Link>
                                    <span> • { this.detail.visit_count } 次浏览 • 最后一次编辑是 { fromNow(this.detail.last_reply_at) } • 来自 { this.comeFrom(this.detail.tab) }</span>
                                </div>
                                {
                                    this.props.store.isLogin
                                    &&
                                    <div className="collection user-select-none">
                                        <button onClick={this.collectionBtn}>{ this.detail.is_collect ? '取消收藏' : '收藏' }</button>
                                    </div>
                                }
                                {
                                    this.props.store.isLogin
                                    &&
                                    this.detail.author_id === this.props.store.userInfo.id
                                    &&
                                    <div className="operation-edit">
                                        <Link to={`/release/${this.detail.id}`}>
                                            <Icon type="edit" title="编辑" />
                                        </Link>
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
                                                    <Link to={`/user/${item.author.loginname}`}>
                                                        <img src={ item.author.avatar_url } alt="头像" />
                                                    </Link>
                                                </div>
                                                <div className="reply-right">
                                                    <div className="reply-author">
                                                        <Link to={`/user/${item.author.loginname}`}>{ item.author.loginname }</Link>
                                                        <span>{ index + 1 }楼•{ fromNow(item.create_at) }</span>
                                                        {
                                                           this.detail.author.loginname == item.author.loginname && <strong>作者</strong>
                                                        }
                                                    </div>
                                                    <div className="operation user-select-none">
                                                        <div>
                                                            <Icon type={item.is_uped ? 'like' : 'like-o'} onClick={this.likeBtn.bind(this, item.id, item.author, index)} />
                                                            <em>{ item.ups.length }</em>
                                                        </div>
                                                        {
                                                            this.props.store.isLogin &&
                                                            <div onClick={this.replyOthers.bind(this, item.author.loginname)}>
                                                                <Icon type="rollback" />
                                                            </div>
                                                        }
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
                    {/* 新建评论 */}
                    <div className="insert-reply" style={{display: !this.props.store.isLogin ? 'none' : ''}}>
                        <div className="tip">添加回复</div>
                        <textarea id="markdown-editor"></textarea>
                        <div className="reply-btn">
                            <button type="button" onClick={this.insertReply}>{ this.insertBtnText }</button>
                        </div>
                    </div>
                </div>
                <Sidebar author={this.detail.author} from="topic" />
            </section>
        )
    }
}



export default Topic;

