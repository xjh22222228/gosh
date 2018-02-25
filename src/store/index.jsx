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
    };
    @observable messageCount = 0;       // 未读消息数
    @observable updateTopicsInfo = {    // 待更新的主题信息
        author: {
            avatar_url: 'https://avatars1.githubusercontent.com/u/15535177?s=460&v=4',
            loginname: '--',
        },
        author_id: '',
        content: '',
        good: false,
        id: '',
        is_collect: false,
        replies: [],
        reply_count: 0,
        tab: '',
        title: '',
        top: false,
        visit_count: 0,
    };

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

    // 将主题详情存入状态以供编辑使用
    @action handleUpdateTopicsInfo (detail) {
        this.updateTopicsInfo = detail;
        window.localStorage.updateTopicsInfo = JSON.stringify(detail);
    }
}


export default Store;