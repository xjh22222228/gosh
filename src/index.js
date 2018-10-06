import React from 'react';
import ReactDOM from 'react-dom';
import App from './App'
import 'antd/dist/antd.css';
import './assets/scss/style.scss'
import './assets/scss/media.scss'
import Store from './store'
import { Provider } from 'mobx-react'
import registerServiceWorker from './registerServiceWorker';

const store = new Store();
ReactDOM.render(
<Provider store={store}>
    <App />
</Provider>, document.getElementById('xiejiahe'));
registerServiceWorker();
