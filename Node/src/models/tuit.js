import mongoose from "mongoose";
const { Schema, model } = mongoose;

const TuitSchema = new Schema({
    nombre: { type: String, required: true },
    comentario: { type: String, required: true },
    fecha: { type: String, required: true },
    hashtags: { type: [String], required: false },
    upvotes: { type: Number, default: true },
    downvotes: { type: Number, required: true }
});

export default model('Tuits', TuitSchema);