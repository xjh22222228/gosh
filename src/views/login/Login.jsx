import React, { Component } from 'react'
import { message } from 'antd';
import './login.scss'
import { observer, inject } from 'mobx-react'
import API_CONFIG from '../../api'
/* eslint-disable */
@inject(stores => stores)
@observer class Login extends Component {

    constructor (props) {
        super(props)
        this.state = {
            accessToken: window.localStorage.save_access_token || '',
        }
    }

    componentWillMount () {
        // 是否存在口令
        if( this.props.store.accessToken ) {
            this.props.history.replace('/');
        }
    }

    /**
     * @param {Number} type 
     * @param {event} e 
     */
    handleChange (e) {
        this.setState({
            accessToken: e.target.value.trim()
        })
    }

    handleSubmit = () => {
        if( !this.state.accessToken ) return message.warning('Access Token不能为空');
        axios.post(API_CONFIG.login, {
            accesstoken: this.state.accessToken
        })
        .then(res => {
            if( res.data.success ) {
                this.props.store.login(this.state.accessToken, res.data);
                this.props.history.replace('/');
            }
        })
        .catch(e => e);
    }

    render () {
        return (
            <section className="login">
                <div className="box">
                    <div className="input last">
                        <input 
                            type="text" 
                            maxLength="50" 
                            value={this.state.accessToken} 
                            placeholder="Access Token" 
                            onChange={this.handleChange.bind(this)} />
                    </div>
                    <div className="get-access-token">
                        <a href="https://cnodejs.org/setting" target="_blank">如何获取Access Token？</a>
                    </div>
                    <div className="submit user-select-none" onClick={this.handleSubmit}>Sign in</div>
                </div>
            </section>
        );
    }
}

export default Login;