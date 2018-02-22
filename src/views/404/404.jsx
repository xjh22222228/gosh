import React from 'react'
import { Link } from 'react-router-dom'
import './404.scss'


var NotMatch = ({ history }) => {
    return (
        <section className="not-match user-select-none">
            <div>
                <img src={require('../../assets/images/not-match.gif')} alt="未匹配" className="pointer-events-none" />
                <div className="page-text">卧槽！页面不见了！</div>
                <div className="go">
                    <Link to="/">首页</Link>
                </div>
            </div>
        </section>
    )
}


export default NotMatch;