
基于React实现的一个有点相似CNode的社区网站, 支持手机端预览

[源码](https://github.com/xjh22222228/React-CNode)

[在线预览](https://xjh22222228.github.io/React-CNode/index.html)


![](https://raw.githubusercontent.com/xjh22222228/React-CNode/master/public/images/11.png)


![](https://raw.githubusercontent.com/xjh22222228/React-CNode/master/public/images/22.png)


### 技术栈
```
"react": "^16.2.0",
"react-dom": "^16.2.0",
"react-router-dom": "^4.2.2",
"mobx": "^3.4.1",
"mobx-react": "^4.3.5",
"antd": "^3.1.6",
"axios": "^0.17.1",
"moment": "^2.20.1",
"webpack": "3.8.1",
"scss": "",
"ES6": "",
"flex": "",
```

### 开发目录
```
.
├── README.md
├── config
│   ├── env.js
│   ├── jest
│   │   ├── cssTransform.js
│   │   └── fileTransform.js
│   ├── paths.js
│   ├── polyfills.js
│   ├── webpack.config.dev.js
│   ├── webpack.config.prod.js
│   └── webpackDevServer.config.js
├── package-lock.json
├── package.json
├── public                       # 静态资源目录
│   ├── images
│   │   ├── 11.png
│   │   └── 22.png
│   ├── index.html
│   └── manifest.json
├── scripts
│   ├── build.js
│   ├── start.js
│   └── test.js
└── src                          # 开发目录，源码文件
    ├── App.jsx                  # views入口文件
    ├── api
    │   └── index.js             # API配置
    ├── assets                   # 资源目录，跟public不同的是assets会被webpack处理
    │   ├── images
    │   │   ├── app-qrcode.png
    │   │   ├── github.svg
    │   │   ├── loading.svg
    │   │   └── not-match.gif
    │   └── scss
    │       ├── _variable.scss
    │       ├── media.scss
    │       └── style.scss
    ├── components               # 存放组件
    │   ├── footer               # 底部组件
    │   │   ├── Footer.jsx
    │   │   └── footer.scss
    │   ├── header               # 头部组件
    │   │   ├── Header.jsx
    │   │   └── header.scss
    │   ├── loading              # 加载组件
    │   │   ├── Loading.jsx
    │   │   └── loading.scss
    │   └── sidebar              # 侧边栏组件
    │       ├── Sidebar.jsx
    │       └── sidebar.scss
    ├── index.js                 # 程序主入口
    ├── registerServiceWorker.js # 这个文件的作用是缓存，下次打开会更快
    ├── router                   # 路由配置
    │   └── index.jsx
    ├── store                    # mobx的状态管理
    │   └── index.jsx
    ├── utils                    # 封装的一些公用方法
    │   └── index.js
    └── views                    # 存放页
        ├── 404                  # 404页
        │   ├── 404.jsx
        │   └── 404.scss
        ├── index                # 首页
        │   ├── Index.jsx
        │   └── style.scss
        ├── login                # 登录页
        │   ├── Login.jsx
        │   └── login.scss
        ├── messages             # 未读消息
        │   ├── Messages.jsx
        │   └── messages.scss
        └── topic                # 主题详情页
            ├── Topic.jsx
            └── style.scss
```


### 小结
+ 本来想打算把所有CNode提供的API开发完成再发布的，然而发现需要处理的东西挺多的，再加上自己并没有太多的时间，一直在拖着。
+ 现在的想法是后续会一个一个把CNode提供的API都加上，不久后就会有一个完整版。
+ CNode主题详情是使用markdown写的，样式可以使用github-markdown-css。
+ CNode提供的API获取主题没有返回总条数导致不能算出总页数, 所以我写死了250页。
+ 大部分功能都需要登录的，所以游客只能进行浏览，如果感兴趣的可以尝试登录。
+ 发布文章和评论使用的markdown编辑器 [simplemde](https://github.com/sparksuite/simplemde-markdown-editor)

...


您的一个👍使得我有继续开源的动力，感谢！


----

``` bash
# install dependencies
npm install

# serve with hot reload at localhost:3888
npm start

# build for production with minification
npm run build

```

----


### License

[MIT](http://opensource.org/licenses/MIT)

