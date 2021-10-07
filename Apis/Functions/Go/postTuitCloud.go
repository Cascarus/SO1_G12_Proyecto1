package cosmos

import (
	"fmt"
	"errors"
	"logs"
	"net/http"
	"time"
	"encoding/json"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)


type Tuit struct {

	Nombre string `json:"nombre"`
	Comentario string `json:"comentario"`
	Fecha string `json:"fecha"`
	Hashtags []string `json:"hashtags"`
	Upvotes int `json:"upvotes"`
	Downvotes int `json:"downvotes"`

}

type Log struct {
    StatusNumber int `json:"statusNumber"`
    Message string `json:"message"`
	Time time.Duration `json:"time"`
}


type Message struct {

	Guardados int `json:"guardados"`
	Api string `json:"api"`
	TiempoCarga string `json:"tiempoCarga"`
	Db string `json:"db"`
}


func PostTuitCloud(w http.ResponseWriter, r *http.Request) {
    t := time.Now()
    
    var newTuit Tuit

    fmt.Println("======================== POSTING TUIT IN SQL CLOUD ========================")

    // Call BindJSON to bind the received JSON to
	if err := json.NewDecoder(r.Body).Decode(&newTuit); err != nil {
		fmt.Fprint(w, "Error reading Tuit!")
		return
	}

    msg, err := Insert(newTuit)

    if err!=nil{
        fmt.Println(err)
        fmt.Fprint(w, Log{StatusNumber:http.StatusInternalServerError, Message:fmt.Sprint(err), Time:time.Since(t) })
    }else{
        fmt.Fprint(w, Log{ StatusNumber:http.StatusCreated, Message:msg, Time:time.Since(t) })
    }

}




func Init() error {

	//dns := os.Getenv("USER")+":"+os.Getenv("PASS")+"@tcp("+os.Getenv("PROXY_ADDRESS")+":1433)/"+os.Getenv("DB") // --> Using proxu
    dns := "root:123456@tcp(34.122.151.115)/Olympics" // --> Using the public IP

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


func Insert(tuit Tuit) (string, error) {

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