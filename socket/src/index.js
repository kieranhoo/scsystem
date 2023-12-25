
import { availableParallelism } from 'node:os';
import cluster from 'node:cluster';
import { setupPrimary } from '@socket.io/cluster-adapter';
import NotificationSocket from './app.js'

import dotenv from 'dotenv'
dotenv.config();

const port = parseInt(process.env.PORT, 10);


if (cluster.isPrimary) {
    const numCPUs = availableParallelism();
    for (let i = 0; i < numCPUs; i++) {
        cluster.fork({
            PORT: port + i
        });
    }

    setupPrimary();
} else {
    const app = new NotificationSocket()
    app.On()
    app.listen(port)
}
