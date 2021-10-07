import cors from 'cors';
import express from 'express';
import morgan from 'morgan';
import dotenv from 'dotenv';

import testsRoutes from '../routes/tests.js';
import subscribe from '../services/pubsub.js'

//import './cloud.js';

import ora from 'ora'

// INITIALIZE =====================================
const app = express();
dotenv.config();

import program from './cloudRT.js'

program()
    .then(() => console.log('Waiting for database events...'))
    .catch(console.error);

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

app.use(testsRoutes)
app.use((req, res, next) => {
    res.status(404).send('404 Not Found');
});
//=================================================


// STATICS ========================================
//=================================================



export default app;