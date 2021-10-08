import { getTuitsCosmos, getTuitsCloud } from "../services/tuits.js";
import { Router } from "express";
const router = Router();

router.get("/getTuitsCosmos", getTuitsCosmos);
router.get("/getTuitsCloud", getTuitsCloud);


export default router;