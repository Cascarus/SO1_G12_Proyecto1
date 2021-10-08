import app from './src/config/app.js';
import { Server } from 'socket.io';
import http from 'http'

import SocketHandler from './src/config/socket.js'

async function main() {

    const server = http.createServer(app);

    const httpServer = server.listen(app.get('port'));
    const io = new Server(httpServer, {
        cors: {
            origin: '*',
        }
    });

    SocketHandler(io)

    await console.log('Servidor publicado en puerto', app.get('port'));

}

main();