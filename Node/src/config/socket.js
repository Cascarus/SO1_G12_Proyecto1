import program from './cloudRT.js'

import { io } from "socket.io-client";

const socket = io("http://localhost:3500");

const SocketHandler = (io) => {

    io.on("connection", (socket) => {
        // console.log(socket.handshake.url);
        console.log("nuevo socket connectado:", socket.id);

        socket.on("hola", () => {
            console.log("Hay alguien");
        });

        socket.on("SQL", (data) => {
            console.log(data)
        });

        socket.on("disconnect", () => {
            console.log(socket.id, "disconnected");
        });

    });

    program(socket)
        .then(() => console.log('Waiting for database events...'))
        .catch(console.error);

}


export default SocketHandler;