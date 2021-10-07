import mysql from 'mysql'
import dotenv from 'dotenv';

dotenv.config();

const pool = mysql.createPool({
    user: process.env.USER, // e.g. 'my-db-user'
    password: process.env.PASS, // e.g. 'my-db-password'
    database: process.env.DB, // e.g. 'my-database'
    host: process.env.DB_ADDR, // e.g. '127.0.0.1'
    // ... Specify additional properties here.
});

pool.getConnection(async function (err, connection) {
    if (err) {
        console.log('Failed connecting to Cloud SQL ');
    }
    else {
        console.log('Connected to CLOUD SQL')
    };

});



export default pool