package main

//nodemon --exec go run main.go --signal SIGTERM

import "github.com/gin-gonic/gin"
import (
	"fmt"
	"net/http"
    "os"
    "time"
    ts "goApi/types"
    cos "goApi/cosmos"
    sql "goApi/cloud"
    ps "goApi/pubSub"
    "github.com/joho/godotenv"
)


func postTuitCosmos(c *gin.Context) {
    t := time.Now()
    var newTuit ts.Tuit

    fmt.Println("======================== POSTING TUIT IN COSMOS ========================")

    // Call BindJSON to bind the received JSON to
    if err := c.BindJSON(&newTuit); err != nil {
        return
    }

    tuits = append(tuits, newTuit)
    msg, err := cos.Create(newTuit)

    if err!=nil{
        fmt.Println(err)
        cosmosLogs = append(cosmosLogs, ts.Log{ StatusNumber:http.StatusInternalServerError, Message:fmt.Sprint(err), Time:time.Since(t) })
        c.JSON(http.StatusInternalServerError, ts.Log{StatusNumber:http.StatusInternalServerError, Message:fmt.Sprint(err), Time:time.Since(t) })
    }else{
        cosmosLogs = append(cosmosLogs, ts.Log{ StatusNumber:http.StatusCreated, Message:msg, Time:time.Since(t) } )
        c.JSON(http.StatusCreated, ts.Log{ StatusNumber:http.StatusCreated, Message:msg, Time:time.Since(t) })
    }

}



func postTuitCloud(c *gin.Context) {
    t := time.Now()
    
    var newTuit ts.Tuit

    fmt.Println("======================== POSTING TUIT IN SQL CLOUD ========================")

    // Call BindJSON to bind the received JSON to
    if err := c.BindJSON(&newTuit); err != nil {
        return
    }

    tuits = append(tuits, newTuit)
    msg, err := sql.Insert(newTuit)

    if err!=nil{
        fmt.Println(err)
        cloudLogs = append(cloudLogs, ts.Log{ StatusNumber:http.StatusInternalServerError, Message:fmt.Sprint(err), Time:time.Since(t) })
        c.JSON(http.StatusInternalServerError, ts.Log{StatusNumber:http.StatusInternalServerError, Message:fmt.Sprint(err), Time:time.Since(t) })
    }else{
        cloudLogs = append(cloudLogs, ts.Log{ StatusNumber:http.StatusCreated, Message:msg, Time:time.Since(t) })
        c.JSON(http.StatusCreated, ts.Log{ StatusNumber:http.StatusCreated, Message:msg, Time:time.Since(t) })
    }

}


func getTuits(c *gin.Context) {
    c.JSON(http.StatusOK, tuits)
}

func getLogsCosmos(c *gin.Context) {
    c.JSON(http.StatusOK, cosmosLogs)
}

func getLogsCloud(c *gin.Context) {
    c.JSON(http.StatusOK, cloudLogs)
}

func validateDataBases(c *gin.Context){
    if ready {
        c.Next()
    }else{
        c.JSON(http.StatusInternalServerError, "Must start a connection")
        c.AbortWithStatus(http.StatusInternalServerError)
    }
}


func pubilshResults(){

    var timeCosmos time.Duration
    var timeCloud time.Duration

    for i := 0; i < len(cosmosLogs); i++ {
        timeCosmos+=cosmosLogs[i].Time
    }
    ps.InitPubSub(ts.Message{ Guardados:len(cosmosLogs), Api:"Go", TiempoCarga: fmt.Sprint(timeCosmos), Db:"Azure Cosmos"})


    for i := 0; i < len(cloudLogs); i++ {
        timeCloud+=cloudLogs[i].Time
    }
    ps.InitPubSub(ts.Message{ Guardados:len(cloudLogs), Api:"Go", TiempoCarga: fmt.Sprint(timeCloud), Db:"Cloud SQL"})

}
/*
	Guardados int `json:"guardados"`
	api string `json:"api"`
	tiempoCarga string `json:"tiempoCarga"`
	db string `json:"db"`
*/

func startLoad(c *gin.Context) {
    
    if !ready {

        _, err := cos.Connect()
        if err != nil {
            ready = false
            c.JSON(http.StatusInternalServerError, "Cosmos DB failed :(")
        }
    
        err1 := sql.Init() 
        if err1 != nil {
            ready = false
            c.JSON(http.StatusInternalServerError, "SQL CLoud failed :(")
        }else{
            ready = true
            c.JSON(http.StatusInternalServerError, "All set!")
        }

    }else{
        c.JSON(http.StatusInternalServerError, "Connection already started")
    }

}


func closeLoad(c *gin.Context){

    ready = false
    c.JSON(http.StatusInternalServerError, "Connections closed")
    
    pubilshResults()

}


var tuits[] ts.Tuit
var cosmosLogs[] ts.Log
var cloudLogs[] ts.Log

var ready bool

func main() {

    if os.Getenv("DB")==""{
        err := godotenv.Load("env.env")
        if err!=nil{
            fmt.Println("Error loading enviroment variables")
        }

    }

    fmt.Println("")
    fmt.Println(" ==========================  SERVIDOR  ========================== ")
    fmt.Println("")
	router := gin.Default()
    router.Use(gin.Recovery()) // Para recuperarse de Errores y enviar un 500

    router.GET("/startLoad/go", startLoad)
    router.POST("/addTuit/cosmos/go", validateDataBases, postTuitCosmos)
    router.POST("/addTuit/cloud/go", validateDataBases, postTuitCloud)
    router.GET("/getTuits/go", getTuits)
    router.GET("/getLogsCosmos/go", getLogsCosmos)
    router.GET("/getLogsCloud/go", getLogsCloud)
    router.GET("/closeLoad/go", closeLoad)

    router.Run()
}

//  /home/sopes1_s2_2021_g14/sopes1/SO1_G12_Proyecto1/Apis/GoApi/