import cors from 'cors';
import express from 'express';
import morgan from 'morgan';
import dotenv from 'dotenv';

import tuitRoutes from '../routes/tuit.js'

import subscribe from '../services/pubsub.js'

//import './cloud.js';
import './cosmos.js'

// INITIALIZE =====================================
const app = express();
dotenv.config();

subscribe(1, process.env.SUBSCRIPTION, 3600)
    .catch(console.error);
//=================================================


// SET PORT =======================================
app.set('port', process.env.PORT);
//=================================================

//MIDDLEWARES =====================================
app.use(express.urlencoded({ extended: false }));
app.use(morgan('method :url :status :res[content-length] - :response-time ms'));
app.use(express.json());
app.use(cors());
//=================================================


// ROUTES =========================================

app.use(function (req, res, next) {
    res.header("Access-Control-Allow-Origin", "*");
    res.header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept");
    next();
});

app.use(tuitRoutes)
app.use((req, res, next) => {
    res.status(404).send('404 Not Found');
});
//=================================================


// STATICS ========================================
//=================================================


// SOCKET ========================================


//=================================================



export default app;