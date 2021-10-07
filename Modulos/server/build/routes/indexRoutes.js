"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const express_1 = require("express");
const generalControler_1 = require("../controllers/generalControler");
class IndexRoutes {
    constructor() {
        this.router = (0, express_1.Router)();
        this.config();
    }
    config() {
        this.router.get('/ram', generalControler_1.indexCtrl.getRam);
        this.router.get('/cpu', generalControler_1.indexCtrl.getCpu);
    }
}
const indexRoutes = new IndexRoutes();
exports.default = indexRoutes.router;
