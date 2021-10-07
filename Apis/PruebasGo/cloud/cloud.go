package cloud

import (
	"log"
    "errors"
    //"time"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Init() error {

	//dns := os.Getenv("USER")+":"+os.Getenv("PASS")+"@tcp("+os.Getenv("PROXY_ADDRESS")+":1433)/"+os.Getenv("DB") // --> Using proxu
    dns := "root:123456@tcp(34.122.151.115)/Olympics" // --> Using the public IP

    /*cfg := mysql.Config{
        User:   "",
        Passwd: "",
        Net:    "tcp",
        Addr:   "",
        DBName: "",
    }*/

	var err error
    //db, err = sql.Open("mysql", cfg.FormatDSN())
	db, err = sql.Open("mysql", dns)
    if err != nil {
        log.Print(err)
		return errors.New("Not connected")
    }
    
    db.SetMaxIdleConns(0)
    db.SetMaxOpenConns(20)

	pingErr := db.Ping()
    if pingErr != nil {
        log.Print(pingErr)
        return errors.New("Not connected")
    }
    return nil
}
