import { getTuits } from "../services/tuits.js";
import { Router } from "express";
const router = Router();

router.get("/getTuits", getTuits);


export default router;