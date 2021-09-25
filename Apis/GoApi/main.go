package main

//nodemon --exec go run main.go --signal SIGTERM

import "github.com/gin-gonic/gin"
import (
	"fmt"
	"net/http"
    "os"
    ts "goApi/types"
    cos "goApi/cosmos"
    sql "goApi/cloud"
    //ps "loadTester/pubSub"
    "github.com/joho/godotenv"
)


func postTuitCosmos(c *gin.Context) {
    
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
        cosmosLogs = append(cosmosLogs, ts.Log{ StatusNumber:http.StatusInternalServerError, Message:fmt.Sprint(err) } )
        c.JSON(http.StatusInternalServerError, ts.Log{StatusNumber:http.StatusInternalServerError, Message:fmt.Sprint(err)})
    }else{
        cosmosLogs = append(cosmosLogs, ts.Log{ StatusNumber:http.StatusCreated, Message:msg } )
        c.JSON(http.StatusCreated, ts.Log{ StatusNumber:http.StatusCreated, Message:msg })
    }

}



func postTuitCloud(c *gin.Context) {
    
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
        cloudLogs = append(cloudLogs, ts.Log{ StatusNumber:http.StatusInternalServerError, Message:fmt.Sprint(err) } )
        c.JSON(http.StatusInternalServerError, ts.Log{StatusNumber:http.StatusInternalServerError, Message:fmt.Sprint(err)})
    }else{
        cloudLogs = append(cloudLogs, ts.Log{ StatusNumber:http.StatusCreated, Message:msg } )
        c.JSON(http.StatusCreated, ts.Log{ StatusNumber:http.StatusCreated, Message:msg })
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

    router.GET("/startLoad", startLoad)
    router.POST("/addTuit/cosmos", validateDataBases, postTuitCosmos)
    router.POST("/addTuit/cloud", validateDataBases, postTuitCloud)
    router.GET("/getTuits", getTuits)
    router.GET("/getLogsCosmos", getLogsCosmos)
    router.GET("/getLogsCloud", getLogsCloud)
    router.GET("/closeLoad", closeLoad)

    router.Run()
}


//   /home/sopes1_s2_2021_g14/sopes1/go/SO1_G12_Proyecto1/Apis/GoApi/