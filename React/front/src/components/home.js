import { Component, Fragment } from "react";

import socketIOClient from 'socket.io-client';

import { FaRedditAlien } from 'react-icons/fa'
import { HiOutlineDocumentReport } from 'react-icons/hi'
import { BiDownvote } from 'react-icons/bi'
import { BiUpvote } from 'react-icons/bi'


const HOST = "http://localhost:1500";
const socket = socketIOClient(HOST);

class Home extends Component {

    constructor(props) {
        super(props)
        this.state = {
            db: "sql"
        }

        this.switchDB = this.switchDB.bind(this)
    }


    componentDidMount() {

        socket.on("SQL", data => {
            console.log(data)
        });

        socket.on("COSMOS", data => {
            console.log(data)
        });

    }


    switchDB() {

        if (this.state.db === "sql") {
            this.setState({ db: "cosmos" })
        }
        else {
            this.setState({ db: "sql" })
        }
    }


    render() {

        return (

            <Fragment>

                <div className="row">

                    <nav className="navbar navbar-expand-lg navbar-light bg-light" style={{ margin: "10px", paddingLeft: "30%", paddingRight: "30%" }}>
                        <div className="container-fluid">
                            <div onClick={() => { this.props.history.push("/") }} style={{ cursor: "pointer" }} className="navbar-brand me-5 negrita">
                                <HiOutlineDocumentReport size={60} ></HiOutlineDocumentReport>
                            </div>

                            <h3 style={{ color: "black" }}> | </h3>

                            <div onClick={() => { this.props.history.push("/") }} style={{ cursor: "pointer" }} className="navbar-brand me-5 negrita">
                                <FaRedditAlien size={60} ></FaRedditAlien>
                            </div>

                            <h3 style={{ color: "black" }}> | </h3>

                            <h4 style={{ color: "black" }} >SQL Cloud</h4>
                            <div className="navbar-brand me-5 negrita">
                                <div className="form-check form-switch">
                                    <input className="form-check-input" type="checkbox" id="flexSwitchCheckDefault" onChange={this.switchDB} />
                                </div>
                            </div>
                            <h3 style={{ color: "black" }} >Cosmo</h3>

                        </div>
                    </nav>

                </div>

                <div className="row">

                    <div className="col-3" style={{ backgroundColor: "#141a15" }} >
                        <h1>
                            Hola
                        </h1>
                    </div>

                    <div className="col-6">

                        <div className="card" style={{ margin: "20px" }}>
                            <div className="card-body bg-dark">

                                <div className="row">
                                    <div className="col">
                                        <h4> Helmut Efrain Najarro Alvarez </h4>
                                    </div>
                                    <div className="col">
                                        <i> 12 de Octube 2021 </i>
                                    </div>
                                </div>
                                <div className="col" style={{ margin: "20px" }}>
                                    <div>
                                        <h3>
                                            jhkahsdkjhasdasdasdas
                                            asdasdasdasdasdasdasd
                                            asdasdasdasdasdasdasd
                                            asdasdasdasdasdasdasd
                                        </h3>
                                    </div>
                                </div>
                                <div className="row">
                                    <div className="col">
                                        <BiDownvote size={30} > </BiDownvote> 45
                                    </div>
                                    <div className="col">
                                        <BiUpvote size={30} > </BiUpvote> 45
                                    </div>
                                </div>
                            </div>
                            <div className="card-footer row" style={{ color: "blue" }}>
                                <div className="col">
                                    <h6> #UNO </h6>
                                </div>
                                <div className="col">
                                    <h6> #UNO </h6>
                                </div>
                            </div>
                        </div>

                    </div>

                    <div className="col-3" style={{ backgroundColor: "#141a15" }}>
                        <h1>
                            Hola
                        </h1>
                    </div>

                </div>

            </Fragment>

        )

    }

}


export default Home