import React from 'react'
import BraftEditor from 'braft-editor'
import 'braft-editor/dist/braft.css'
import { Link } from 'react-router-dom'
import { message } from 'antd'
import axios from 'axios'
import API_CONFIG from '../../api'
import './style.scss'

class Editor extends React.Component{
    state = { tabValue:'',topic_name:'',topic_content:null ,topic_content_html:null}
    onSelectValue = e => {
        this.setState({
            tabValue:e.target.value
        })
    }
    onTopicTitle = e => (
        this.setState({
            topic_name:e.target.value
        })
    )
    handleChange = (content) => {
        this.setState({
            topic_content:content
        })
    }
    
    handleHTMLChange = (html) => {
        this.setState({
            topic_content_html:html
        })
    }

    onSubmit = () => {
        if(this.state.tabValue === ''){
            message.error('必须选择一个板块！')
        }else{
            if(this.state.topic_name === ''){
                message.error('标题不能为空')
            }else{
                if(this.state.topic_content === null && this.state.topic_content_html === null){
                    message.error('文章内容不能为空')
                }else{
                    axios.post(API_CONFIG.postTopic,{
                        title:this.state.topic_name,
                        tab:this.state.tabValue,
                        content:this.state.topic_content_html
                    })
                    .then(
                        res => (
                            this.props.history.replace('/')
                        )
                    )
                    .catch(e => e)
                }
            }
        }
    }
    render(){
        const editorProps = {
            width:700,
            height: 500,
            initialContent: this.state.topic_content,
            onChange: this.handleChange,
            onHTMLChange: this.handleHTMLChange
        }
        return(
                <div className="create_edtior">
                    <div className="create_edtior_header">
                        <ul>
                            <li><Link to="/">主页</Link></li>
                            <li><em>/</em></li>
                            <li>发布话题</li>
                        </ul>
                    </div>
                    <div className="create_quill">
                        <div className="topic_select">
                            <label>选择版块：</label>
                            <select name="tab" value={this.state.tabValue} onChange={this.onSelectValue}>
                                <option value="">请选择</option>                            
                                <option value="ask">问答</option>
                                <option value="share">分享</option>
                                <option value="job">招聘</option>
                                <option value="good">精华</option>
                                <option value="dev">客户端测试</option>
                            </select>
                        </div>
                        <div className="create_topic">
                            <div className="topic_title">
                                <input type="text" name="topic_name" placeholder="标题10个字以上" onChange={this.onTopicTitle}/>
                            </div>
                            <BraftEditor {...editorProps}/>
                        </div>
                        <div className="submit">
                            <button onClick={this.onSubmit}>提交</button>
                        </div>
                    </div>
                </div>
        )
    }
}

export default class Publish extends React.Component{
    render(){
        return(
            <div className="create_content">
                <Editor/>
                <div className="create_prompt">

                </div>
            </div>
        )
    }
}