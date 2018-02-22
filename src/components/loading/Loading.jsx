import React, { Component } from 'react'
import PropTypes from 'prop-types'
import './loading.scss'


class Loading extends Component {

    static defaultProps = {
        loading: false
    }

    render () {
        var style = this.props.loading ? {} : {display: 'none'};
        return (
            <div className="loading" style={style}>
                <img src={require('../../assets/images/loading.svg')} alt="loading" />
            </div>
        );
    }
}

Loading.propTypes = {
    loading: PropTypes.bool
}


export default Loading;