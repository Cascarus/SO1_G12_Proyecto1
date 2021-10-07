import tuitSchema from '../models/tuit.js'

const changeStream = tuitSchema.watch(
    [
        { $match: { "operationType": { $in: ["insert", "update", "replace"] } } },
        { $project: { "_id": 1, "fullDocument": 1, "ns": 0, "documentKey": 0 } }
    ],
    { fullDocument: "updateLookup" });

changeStream.on('change', (data) => {
    console.log(data); // You could parse out the needed info and send only that data. 
});

export const getTuits = async (req, res) => {
    const tuits = await tuitSchema.find({});
    res.send(tuits);
}
