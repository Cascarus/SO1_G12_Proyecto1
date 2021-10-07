import { Router } from "express";
import { indexCtrl } from '../controllers/generalControler'

class IndexRoutes{
    public router: Router = Router();

    constructor(){
        this.config();
    }

    config(): void{
        this.router.get('/ram', indexCtrl.getRam );
        this.router.get('/cpu', indexCtrl.getCpu);
    }
}

const indexRoutes = new IndexRoutes();
export default indexRoutes.router;