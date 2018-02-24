import React , {Component}from 'react'
import { HashRouter as Router, Switch, Route} from 'react-router-dom'
import HomePage from '../views/index/Index'
import Login from '../views/login/Login'
import NotMatch from '../views/404/404'
import Header from '../components/header/Header'
import Footer from '../components/footer/Footer'
import Topic from '../views/topic/Topic'
import Messages from '../views/messages/Messages'
import Publish from '../views/publish/Publish'

class Routes extends Component {

    render () {
        return (
            <Router>
                <div className="router-view">
                    <Header />
                        <Switch>
                            <Route path="/" exact component={HomePage} />
                            <Route path="/login" component={Login} />
                            <Route path="/topic/publish" component={Publish} />
                            <Route path="/topic/:id" component={Topic} />
                            <Route component={NotMatch} />
                        </Switch>
                    <Footer />
                </div>
            </Router>
        );
    }
}

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
                        <Route component={NotMatch} />
                    </Switch>
                <Footer />
            </div>
        </Router>
    )
};

export default Routes;