import app from './src/config/app.js';

async function main(){
    app.listen(app.get('port'));
    await console.log('Servidor publicado en puerto', app.get('port'));
}

main();