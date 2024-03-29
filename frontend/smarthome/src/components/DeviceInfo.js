import React, {Component} from 'react';
import { Link } from 'react-router-dom';

import '../slate/bootstrap.min.css';


class DeviceInfo extends Component {
    constructor(props) {
        super(props);
        this.state = {
            deviceinfo: null,
            // devices {
            //     name  : ""
            //     ip    : ""
            //     alias : ""
            //     state : ""
            // }
        }
    }
    componentDidMount() {
        const { ip } = this.props.match.params
        const fetchUrl = process.env.REACT_APP_BACKEND_URL + "/devices/" + ip
        console.log("GET" , fetchUrl)
        fetch(fetchUrl)
        .then(res => res.json())
        .then((data) => {
            this.setState({ deviceinfo: data })
        })
        .catch(console.log)
    }
    render() {
        const {deviceinfo} = this.state;
        if (deviceinfo) {
            return (
                <div>
                    <h2>Device Info for: {this.props.match.params.ip}</h2>
                    <div><pre>{JSON.stringify(deviceinfo, null, '    ') }</pre></div>
                    <Link to="/devices">
                        <button className="btn btn-primary">Back</button>
                    </Link>
                </div>
            )
        }
        return (
            <div>
                <h2>No Device Info for: {this.props.match.params.ip}</h2>
            </div>
        )
    }
}

export default DeviceInfo;