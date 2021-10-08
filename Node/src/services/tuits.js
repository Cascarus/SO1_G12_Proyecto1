import tuitSchema from '../models/tuit.js'

export const getTuits = async (req, res) => {
    const tuits = await tuitSchema.find({});
    res.send(tuits);
}
