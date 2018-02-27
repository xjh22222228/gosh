import  { observable, action } from 'mobx'
import API_CONFIG from '../api'

/* eslint-disable */
class Store {

    @observable accessToken = window.localStorage.access_token || '';
    @observable isLogin = false;
    @observable userInfo = {
        avatar_url: 'data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAMAAAACAQMAAACnuvRZAAAAA1BMVEX29vYACyOqAAAACklEQVQI12MAAgAABAABINItbwAAAABJRU5ErkJggg==',     // 头像
        id: '',             // 用户 id
        loginname: '',      // 用户名
    };
    @observable messageCount = 0;       // 未读消息数

    /**
     * @func 登录
     * @param {String} accessToken 
     * @param {Object} userInfo 
     */
    @action login (accessToken, userInfo) {
        window.localStorage.access_token = window.localStorage.save_access_token = this.accessToken = accessToken;
        this.userInfo = userInfo;
        this.isLogin = true;
        this.fetchMessageCount();
    }

    @action logout () {
        window.localStorage.removeItem('access_token');
        this.accessToken = '';
        this.isLogin = false;
    }
    // 检查是否已登录
    @action checkLogin () {
        if( this.accessToken ) {
            axios.post(API_CONFIG.login)
            .then(res => {
                if( res.data.success ) {
                    this.isLogin = true;
                    this.userInfo = res.data;
                    this.fetchMessageCount();
                }
            })
            .catch(e => e);
        }
    }

    // 获取未读消息数
    @action fetchMessageCount () {
        axios.get(API_CONFIG.messageCount)
        .then(res => {
            if( res.data.success ) {
                this.messageCount = res.data.data;
            }
        })
        .catch(e => e);
    }
}


export default Store;