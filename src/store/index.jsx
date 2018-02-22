import  { observable, action } from 'mobx'
import API_CONFIG from '../api'

/* eslint-disable */
class Store {

    @observable accessToken = window.localStorage.access_token || '';
    @observable isLogin = false;
    @observable userInfo = {
        avatar_url: '',     // 头像
        id: '',             // 用户 id
        loginname: '',      // 用户名
    }

    /**
     * @func 登录
     * @param {String} accessToken 
     * @param {Object} userInfo 
     */
    @action login (accessToken, userInfo) {
        window.localStorage.access_token = window.localStorage.save_access_token = this.accessToken = accessToken;
        this.userInfo = userInfo;
        this.isLogin = true;
    }

    @action logout (a) {
        window.localStorage.removeItem('access_token');
        this.accessToken = '';
        this.isLogin = false;
    }

    @action checkLogin () {
        if( this.accessToken ) {
            axios.post(API_CONFIG.login)
            .then(res => {
                if( res.data.success ) {
                    this.isLogin = true;
                    this.userInfo = res.data;
                }
            })
            .catch(e => e);
        }
    }
}


export default Store;