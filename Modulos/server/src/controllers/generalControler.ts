import { Request, Response } from "express";
import * as fs from 'fs';

class IndexCtrl{

    public async getRam(req: Request, res:Response){
        //let ram = fs.readFileSync('/elements/procs/rammod').toString().split(',')
        let ram = fs.readFileSync('/proc/rammod').toString().split(',');
        let tempRam = {
            ram_total: ram[0],
            ram_uso: ram[1],
            ram_libre: ram[2],
            ram_percent: ram[3]
        }
        console.log(tempRam);
        res.json(tempRam);
    }

    public async getCpu(req: Request, res:Response){
        //let CPU = fs.readFileSync('/elements/procs/processlistmod').toString().split(',')
        let CPU = fs.readFileSync('/proc/processlistmod').toString().split(',')
        let tempCpu = {
            procesos: CPU[0],
            cpu_uso: CPU[1]
        }
        console.log(tempCpu);
        res.json(tempCpu);
    }

}

export const indexCtrl = new IndexCtrl();