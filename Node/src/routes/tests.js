import { Router } from "express";
const router = Router();

import {holaMundo} from '../services/tests.js'

router.get('/hola', holaMundo)


export default router