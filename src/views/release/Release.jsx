import React, { Component } from 'react'
import API_CONFIG from '../../api'
import { observer, inject } from 'mobx-react'
import { Link } from 'react-router-dom'
import { Select, Input, message } from 'antd'
import SimpleMDE from 'simplemde'
import './release.scss'
import { observable } from 'mobx'

/* eslint-disable */
@inject(stores => stores)
@observer class Release extends Component {

    @observable currentStatus = true;   // true = 发布话题， false 编辑话题

    constructor (props) {
        super(props);
        this.state = {
            title: '',
            tab: 'dev',
        };
    }

    handleChangeSelect = val => {
        this.setState({
            tab: val
        });
    }

    handleChangeTitle = e => {
        this.setState({
            title: e.target.value.trim()
        });
    }

    // 初始化markdown编辑器
    initMarkdownEditor = () => {
        this.simplemde = new SimpleMDE({
            element: document.getElementById("markdown-editor"),
            spellChecker: false, 				// 启用拼写检查，会有背景色
            autoDownloadFontAwesome: false,		// 是否需要下载字体图标
        });
    }

    // 发布或更新话题
    releaseTopics = () => {
        try {
            if( !this.props.store.isLogin ) throw new Error('请登录后再提交');
            if( this.state.title.length < 10 ) throw new Error('标题至少10字以上');
            var val = this.simplemde.value();
            if( !val ) throw new Error('主体内容不能为空');
        } catch (e) {
            return message.warning(e.message);
        }
        var url = this.currentStatus ? API_CONFIG.newTopics : API_CONFIG.updateTopics;
        axios.post(url, {
            title: this.state.title,
            tab: this.state.tab,
            content: val,
            topic_id: this.currentStatus ? undefined : this.props.match.params.id,
        })
        .then(res => {
            if( res.data.success ) {
                this.currentStatus ? message.success('发布成功') : message.success('更新成功');
                this.props.history.push(`/topic/${res.data.topic_id}`);
            }
        })
        .catch(e => e);
    }

    componentDidMount () {
        this.initMarkdownEditor();
        // 如果是编辑状态
        if( this.props.match.params.id !== 'create' ) {
            this.currentStatus = false;
            axios.get(`${API_CONFIG.topicDetail}${this.props.match.params.id}`, {
                params: {
                    mdrender: false,
                }
            })
            .then(res => {
                if( res.data.success ) {
                    this.setState({
                        title: res.data.data.title,
                        tab: res.data.data.tab,
                    });
                    this.simplemde.value(res.data.data.content);
                }
            })
            .catch(e => {
                message.warning('不存在此话题');
                this.props.history.replace('/');
            });
        }
    }

    render () {
        const Option = Select.Option;
        return (
            <section className="release index-section">
                <div className="topics-container release-left">
                    <div className="top">
                        <Link to="/">主页</Link>
                        <em> / </em>
                        <span>{ this.currentStatus ? '发布话题' : '编辑话题'}</span>
                    </div>
                    <div className="select-tab">
                        <span>选择版块：</span>
                        <Select 
                            defaultValue={this.state.tab} 
                            style={{ width: 200 }} 
                            onChange={this.handleChangeSelect} 
                            value={this.state.tab}>
                            <Option value="dev">客户端测试</Option>
                            <Option value="share">分享</Option>
                            <Option value="ask">问答</Option>
                            <Option value="job">招聘</Option>
                        </Select>
                    </div>
                    <div className="title">
                        <Input 
                            placeholder="标题字数 10字以上" 
                            onChange={this.handleChangeTitle} 
                            value={this.state.title} />
                    </div>
                    <div className="editor">
                        <textarea id="markdown-editor"></textarea>
                        <div className="release-btn">
                            <button onClick={this.releaseTopics}>{ this.currentStatus ? '发布' : '更新'}</button>
                        </div>
                    </div>
                </div>
                {/* 侧边栏 */}
                <div className="release-sidebar">
                    <div className="block-box">
                        <div className="title-top">Markdown 语法参考</div>
                        <div className="inner">
                            <p>### 单行的标题</p>
                            <p>**粗体**</p>
                            <p>`console.log('行内代码')`</p>
                            <p>```js\n code \n``` 标记代码块</p>
                            <p>[内容](链接)</p>
                            <p>![文字说明](图片链接)</p>
                            <p>
                                <a href="https://segmentfault.com/markdown" target="_blank">Markdown 文档</a>
                            </p>
                        </div>
                    </div>
                    <div className="block-box">
                        <div className="title-top">话题发布指南</div>
                        <div className="inner">
                            <p>尽量把话题要点浓缩到标题里</p>
                            <p>代码含义和报错可在 <a href="https://segmentfault.com/t/node.js" target="_blank">SegmentFault</a> 提问</p>
                        </div>
                    </div>
                </div>
            </section>
        );
    }
}


export default Release;