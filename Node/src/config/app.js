import cors from 'cors';
import express from 'express';
import morgan from 'morgan';
import dotenv from 'dotenv';

import testsRoutes from '../routes/tests.js';

// INITIALIZE =====================================
const app = express();
dotenv.config();
import subscribe from '../services/pubsub.js'

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
app.use((req, res, next) =>{
    res.status(404).send('404 Not Found');
});
//=================================================


// STATICS ========================================
//=================================================



export default app;