"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    Object.defineProperty(o, k2, { enumerable: true, get: function() { return m[k]; } });
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || function (mod) {
    if (mod && mod.__esModule) return mod;
    var result = {};
    if (mod != null) for (var k in mod) if (k !== "default" && Object.prototype.hasOwnProperty.call(mod, k)) __createBinding(result, mod, k);
    __setModuleDefault(result, mod);
    return result;
};
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.indexCtrl = void 0;
const fs = __importStar(require("fs"));
class IndexCtrl {
    getRam(req, res) {
        return __awaiter(this, void 0, void 0, function* () {
            //let ram = fs.readFileSync('/elements/procs/rammod').toString().split(',')
            let ram = fs.readFileSync('/proc/rammod').toString().split(',');
            let tempRam = {
                ram_total: ram[0],
                ram_uso: ram[1],
                ram_libre: ram[2],
                ram_percent: ram[3]
            };
            console.log(tempRam);
            res.json(tempRam);
        });
    }
    getCpu(req, res) {
        return __awaiter(this, void 0, void 0, function* () {
            //let CPU = fs.readFileSync('/elements/procs/processlistmod').toString().split(',')
            let CPU = fs.readFileSync('/proc/processlistmod').toString().split(',');
            let tempCpu = {
                procesos: CPU[0],
                cpu_uso: CPU[1]
            };
            console.log(tempCpu);
            res.json(tempCpu);
        });
    }
}
exports.indexCtrl = new IndexCtrl();
