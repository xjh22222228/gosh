import React, { Component } from 'react'
import { HashRouter as Router, Switch, Route} from 'react-router-dom'
import HomePage from '../views/index/Index'
import Login from '../views/login/Login'
import NotMatch from '../views/404/404'
import Header from '../components/header/Header'
import Footer from '../components/footer/Footer'
import Topic from '../views/topic/Topic'

class Routes extends Component {

    render () {
        return (
            <Router>
                <div className="router-view">
                    <Header />
                        <Switch>
                            <Route path="/" exact component={HomePage} />
                            <Route path="/login" component={Login} />
                            <Route path="/topic/:id" component={Topic} />
                            <Route component={NotMatch} />
                        </Switch>
                    <Footer />
                </div>
            </Router>
        );
    }
}


export default Routes;