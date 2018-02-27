import React from 'react'
import { HashRouter as Router, Switch, Route} from 'react-router-dom'
import Header from '../components/header/Header'
import Footer from '../components/footer/Footer'
import HomePage from '../views/index/Index'
import Login from '../views/login/Login'
import NotMatch from '../views/404/404'
import Topic from '../views/topic/Topic'
import Messages from '../views/messages/Messages'
import Release from '../views/release/Release'
import User from '../views/user/User'
import Collections from '../views/collections/Collections'


const Routes = () => {
    return (
        <Router>
            <div className="router-view">
                <Header />
                    <Switch>
                        <Route path="/" exact component={HomePage} />
                        <Route path="/login" component={Login} />
                        <Route path="/topic/:id" component={Topic} />
                        <Route path="/messages/" component={Messages} />
                        <Route path="/Release/:id" component={Release} />
                        <Route path="/user/:loginname" exact component={User} />
                        <Route path="/user/:loginname/collections" component={Collections} />
                        <Route component={NotMatch} />
                    </Switch>
                <Footer />
            </div>
        </Router>
    )
};

export default Routes;