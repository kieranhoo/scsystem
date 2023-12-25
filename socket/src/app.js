import express from 'express';
import { createServer } from 'node:http';
import { Server } from 'socket.io';
import { createAdapter } from '@socket.io/cluster-adapter';

export default class NotificationSocket {
    constructor() {
        this.app = express();
        this.server = createServer(this.app)
        this.io = new Server(this.server, {
            connectionStateRecovery: {},
            adapter: createAdapter()
        });
        this.channel = "notifications"
    }

    On() {
        this.io.on('connection', async (socket) => {
            console.log("connected")
            socket.on(this.channel, async (msg) => {
                this.io.emit(this.channel, msg);
            });
        });
    }

    listen(port) {
        this.server.listen(port, () => {
            console.log(`server running at http://localhost:${port}`);
        });
    }
}