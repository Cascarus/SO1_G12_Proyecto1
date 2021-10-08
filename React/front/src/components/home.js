import { Component, Fragment } from "react";

import socketIOClient from 'socket.io-client';

import { FaRedditAlien } from 'react-icons/fa'
import { HiOutlineDocumentReport } from 'react-icons/hi'
import { BiDownvote } from 'react-icons/bi'
import { BiUpvote } from 'react-icons/bi'
import axios from 'axios';


const HOST = "http://35.223.137.189:1500";
const socket = socketIOClient(HOST);

var sqlT = [];
var cosmosT = [];

class Home extends Component {

    constructor(props) {
        super(props)
        this.state = {
            db: "sql",
            tuits: []
        }

        this.switchDB = this.switchDB.bind(this)
        this.showHastags = this.showHastags.bind(this)
    }


    async componentDidMount() {

        await axios.get(`${HOST}/getTuitsCloud`)
            .then((res) => {
                this.setState({ tuits: res.data })
                sqlT = res.data
            }).catch((err) => {
                console.log(err)
            })

        socket.on("SQL", data => {
            sqlT.unshift(data)
            this.setState({ tuits: sqlT })
        });

        socket.on("COSMOS", data => {
            cosmosT.unshift(data.fullDocument)
            this.setState({ db: "cosmos", tuits: cosmosT })
        });

    }


    showHastags(hashtags) {

        if (this.state.db === "sql") {
            return (
                hashtags.split(',').map((h) => {
                    return (
                        <div className="col">
                            <h4> #{h} </h4>
                        </div>
                    )
                })
            )
        } else {
            return (
                hashtags.map((h) => {
                    return (
                        <div className="col">
                            <h4> #{h} </h4>
                        </div>
                    )
                })
            )
        }

    }

    async switchDB() {

        if (this.state.db === "sql") {

            await axios.get(`${HOST}/getTuitsCosmos`)
                .then((res) => {
                    this.setState({ tuits: res.data, db: "cosmos" })
                    cosmosT = res.data
                }).catch((err) => {
                    console.log(err)
                })
        }
        else {

            await axios.get(`${HOST}/getTuitsCloud`)
                .then((res) => {
                    this.setState({ tuits: res.data, db: "sql" })
                    sqlT = res.data
                }).catch((err) => {
                    console.log(err)
                })
        }

    }


    render() {

        if (this.state.tuits.length > 0) {
            return (

                <Fragment>

                    <div className="row">

                        <nav className="navbar navbar-expand-lg navbar-light bg-light" style={{ margin: "10px", paddingLeft: "30%", paddingRight: "30%" }}>
                            <div className="container-fluid">
                                <div onClick={() => { this.props.history.push("/dashboard") }} style={{ cursor: "pointer" }} className="navbar-brand me-5 negrita">
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
                                <h3 style={{ color: "black" }} >Cosmos</h3>

                            </div>
                        </nav>

                    </div>

                    <div className="row">

                        <div className="col-3" style={{ backgroundColor: "#141a15" }} >
                        </div>

                        <div className="col-6">


                            {this.state.tuits.map((t, i) => {
                                return (

                                    <div className="card" style={{ margin: "30px" }} key={i + t.nombre}>
                                        <div className="card-body bg-dark">

                                            <div className="row">
                                                <div className="col">
                                                    <h4> {t.nombre} </h4>
                                                </div>
                                                <div className="col">
                                                    <i> {t.fecha} </i>
                                                </div>
                                            </div>
                                            <div className="col" style={{ margin: "20px" }}>
                                                <div>
                                                    <h3>
                                                        {t.comentario}
                                                    </h3>
                                                </div>
                                            </div>
                                            <div className="row">
                                                <div className="col">
                                                    <BiDownvote size={30} > </BiDownvote> {t.downvotes}
                                                </div>
                                                <div className="col">
                                                    <BiUpvote size={30} > </BiUpvote> {t.upvotes}
                                                </div>
                                            </div>
                                        </div>
                                        <div className="card-footer row" style={{ color: "blue" }}>
                                            {
                                                this.showHastags(t.hashtags)
                                            }
                                        </div>
                                    </div>

                                )
                            })}

                        </div>

                        <div className="col-3" style={{ backgroundColor: "#141a15" }}>
                        </div>

                    </div>

                </Fragment>

            )
        } else {
            return (
                <div>

                    <div className="row">

                        <nav className="navbar navbar-expand-lg navbar-light bg-light" style={{ margin: "10px", paddingLeft: "30%", paddingRight: "30%" }}>
                            <div className="container-fluid">
                                <div onClick={() => { this.props.history.push("/dashboard") }} style={{ cursor: "pointer" }} className="navbar-brand me-5 negrita">
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
                                <h3 style={{ color: "black" }} >Cosmos</h3>

                            </div>
                        </nav>

                    </div>

                    <h1> Nada por aqui ...</h1>
                </div>
            )
        }

    }

}


export default Home