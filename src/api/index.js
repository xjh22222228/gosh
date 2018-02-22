const API_CONFIG = {
    // 获取用户详情信息
    login: 'https://cnodejs.org/api/v1/accesstoken',
    // 获取主题列表
    topics: 'https://cnodejs.org/api/v1/topics',
    // 获取主题详情包括回复 GET
    topicDetail: 'https://cnodejs.org/api/v1/topic/',
    // 收藏主题 POST
    collection: 'https://cnodejs.org/api/v1/topic_collect/collect',
    // 取消收藏主题 POST
    cancelCollection: 'https://cnodejs.org/api/v1/topic_collect/de_collect',
    // 点赞/取消点赞 POST  /reply/:reply_id/ups 为评论点赞
    like: 'https://cnodejs.org/api/v1/reply/',
};


export default API_CONFIG;