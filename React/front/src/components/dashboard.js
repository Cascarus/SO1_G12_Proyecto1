import { Component, Fragment } from "react";

import socketIOClient from 'socket.io-client';
import { FaRedditAlien } from 'react-icons/fa'
import { HiOutlineDocumentReport } from 'react-icons/hi'

import { BiDownvote } from 'react-icons/bi'
import { BiUpvote } from 'react-icons/bi'
import { ImNewspaper } from 'react-icons/im'


import axios from 'axios';

import Grafica from "./grafica";

const HOST = "http://35.223.137.189:1500";
const socket = socketIOClient(HOST);

var sqlT = [];
var cosmosT = [];


class Dashboard extends Component {

    constructor(props) {
        super(props);

        this.state = {
            db: "sql",
            tuits: [],
            downVotes: [],
            upVotes: []

        }

        this.switchDB = this.switchDB.bind(this)
    }


    async componentDidMount() {

        await axios.get(`${HOST}/getTuitsCloud`)
            .then((res) => {
                this.setState({ tuits: res.data })
                sqlT = res.data
                this.setVotes(res.data)


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

    setVotes(res) {

        var upV = res.map(t => (
            t.upvotes
        ))
        var downV = res.map(t => (
            t.downvotes
        ))
        this.setState({ upVotes: upV, downVotes: downV })
        console.log(this.state)
    }

    async switchDB() {

        if (this.state.db === "sql") {

            await axios.get(`${HOST}/getTuitsCosmos`)
                .then((res) => {
                    console.log(res.data)
                    this.setState({ tuits: res.data, db: "cosmos" })
                    cosmosT = res.data
                    this.setVotes(res.data)
                }).catch((err) => {
                    console.log(err)
                })
        }
        else {

            await axios.get(`${HOST}/getTuitsCloud`)
                .then((res) => {
                    this.setState({ tuits: res.data, db: "sql" })
                    sqlT = res.data
                    this.setVotes(res.data)
                }).catch((err) => {
                    console.log(err)
                })
        }

    }

    render() {

        if (this.state.upVotes.length > 0) {
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
                                <h3 style={{ color: "black" }} >Cosmo</h3>

                            </div>
                        </nav>

                    </div>

                    <div className="row" style={{ padding: "30px" }}>

                        <div className="col">
                            <div className="card card-body bg-success">
                                <h3>
                                    <BiUpvote size={30} > </BiUpvote> {this.state.upVotes.length} Upvotes
                                </h3>
                            </div>
                        </div>

                        <div className="col">
                            <div className="card card-body bg-primary">
                                <h3>
                                    <ImNewspaper size={30} ></ImNewspaper> {this.state.tuits.length} Noticias
                                </h3>
                            </div>
                        </div>

                        <div className="col">
                            <div className="card card-body bg-warning">
                                <h3 style={{ color: "black" }}>
                                    <BiDownvote size={30}> </BiDownvote> {this.state.downVotes.length} Downvotes
                                </h3>
                            </div>
                        </div>

                    </div>


                    <div className="row" style={{ justifyContent: "center", display: "flex", flexWrap: "wrap" }}>

                        <div className="col" style={{ justifyContent: "flex-end", display: "flex", flexWrap: "wrap" }}>
                            <div className="card">
                                <div className="card-header">
                                    <h2 style={{ color: "black" }} > Upvotes </h2>
                                </div>
                                <div className="card-body bg-dark">
                                    <Grafica data={this.state.upVotes} height={500} width={600} max={1000} key={1} ></Grafica>
                                </div>
                            </div>
                        </div>
                        <div className="col" style={{ justifyContent: "flex-start", display: "flex", flexWrap: "wrap" }}>
                            <div className="card">
                                <div className="card-header">
                                    <h2 style={{ color: "black" }} > Upvotes </h2>
                                </div>
                                <div className="card-body bg-dark">
                                    <Grafica data={this.state.downVotes} height={500} width={600} max={1000} key={2} ></Grafica>
                                </div>
                            </div>
                        </div>

                    </div>


                </Fragment>

            )
        }
        else {
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
                                <h3 style={{ color: "black" }} >Cosmo</h3>

                            </div>
                        </nav>

                    </div>

                </Fragment>

            )
        }

    }

}


export default Dashboard