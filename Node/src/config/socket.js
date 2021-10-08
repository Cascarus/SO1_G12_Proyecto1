import cloud from './cloudRT.js'
import cosmos from './cosmosRT.js'

import { io } from "socket.io-client";

const socket = io("http://localhost:3500");

const SocketHandler = (io) => {

    io.on("connection", (socket) => {
        // console.log(socket.handshake.url);
        console.log("nuevo socket connectado:", socket.id);


        socket.on("SQL", (data) => {
            console.log(data)
            io.emit("SQL", data)
        });

        socket.on("COSMOS", (data) => {
            console.log(data.fullDocument)

            io.emit("COSMOS", data)
        });

        socket.on("disconnect", () => {
            console.log(socket.id, "disconnected");
        });

    });

    cloud(socket)
        .then(() => console.log('Waiting for database events...'))
        .catch(console.error);

    cosmos(socket)
}


export default SocketHandler;