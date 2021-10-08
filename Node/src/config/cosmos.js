import mongoose from "mongoose";
import dotenv from 'dotenv';

dotenv.config();

(async () => {
    try {
        const db = await mongoose.connect(process.env.MONGO_CS, {
            useNewUrlParser: true,
            useUnifiedTopology: true
        });
        console.log("Mongodb is connected to", db.connection.host);


    } catch (error) {
        console.error("Mongodb connection failed", error);
    }
})();
