import MySQLEvents from '@rodrigogs/mysql-events';
import pool from './cloud.js'

const program = async () => {


    const instance = new MySQLEvents(pool, {
        startAtEnd: true
    });

    await instance.start();

    instance.addTrigger({
        name: 'OLIMPIC',
        expression: '*',
        statement: MySQLEvents.STATEMENTS.INSERT,
        onEvent: (event) => { // You will receive the events here
            console.log(event.affectedRows[0].after);
        },
    });

    instance.on(MySQLEvents.EVENTS.CONNECTION_ERROR, console.error);
    instance.on(MySQLEvents.EVENTS.ZONGJI_ERROR, console.error);
}

export default program;