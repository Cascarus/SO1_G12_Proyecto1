import tuitSchema from '../models/tuit.js'
import pool from '../config/cloud.js'


export const getTuitsCosmos = async (req, res) => {
    const tuits = await tuitSchema.find({});
    res.send(tuits);
}


export const getTuitsCloud = async (req, res) => {

    pool.query('select * from OLIMPIC order by id desc', function (error, results, fields) {
        if (error) res.send({ status: 500, msg: error });
        res.send(results)
    });

}
