package cloud

import (
	"log"
    "os"
    "fmt"
    "errors"
    //"time"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
    ts "goApi/types"
)

var db *sql.DB

func Init() error {

	//dns := os.Getenv("USER")+":"+os.Getenv("PASS")+"@tcp("+os.Getenv("PROXY_ADDRESS")+":1433)/"+os.Getenv("DB") // --> Using proxu
    dns := os.Getenv("USER")+":"+os.Getenv("PASS")+"@tcp("+os.Getenv("DB_ADDR")+")/"+os.Getenv("DB") // --> Using the public IP

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

    fmt.Println("asdas")

	pingErr := db.Ping()
    if pingErr != nil {
        log.Print(pingErr)
        return errors.New("Not connected")
    }
    return nil
}


func Insert(tuit ts.Tuit) (string, error) {

    err:=Init()
    if err != nil {
        return "", err
    }
    

    hashtags := ""
    for i := 0; i < len(tuit.Hashtags); i++ {
        hashtags+=tuit.Hashtags[i]+","
    }

    query := "INSERT INTO OLIMPIC (nombre, comentario, fecha, hashtags, upvotes, downvotes) VALUES (?, ?, STR_TO_DATE(?, '%d/%m/%Y'), ?, ?, ?)"

    result, err := db.Exec(query, tuit.Nombre, tuit.Comentario,tuit.Fecha, hashtags, tuit.Upvotes, tuit.Downvotes)
    if err != nil {
        fmt.Print(err)
        return "", errors.New("Couldn't load tuit")
    }
    id, err := result.LastInsertId()
    if err != nil {
        fmt.Print(err)
        return ""+fmt.Sprint(id), errors.New("Couldn't load tuit")
    }

    return "Tuit loaded successfully with id: "+fmt.Sprint(id), nil
}

/*
use Olympics;

create table OLIMPIC(

	idOlimpic INT AUTO_INCREMENT PRIMARY KEY,
    nombre VARCHAR(50),
    comentario VARCHAR(300),
    fecha DATE,
    hashtags VARCHAR(300),
    upvotes INT,
    downvotes INT

);

select * from OLIMPIC;
*/