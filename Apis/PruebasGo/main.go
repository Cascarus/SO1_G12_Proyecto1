package main

//nodemon --exec go run main.go --signal SIGTERM

import "github.com/gin-gonic/gin"
import (
	"fmt"
	"net/http"
    sql "goApi/cloud"
)


func postTuitCloud(c *gin.Context) {
    
    fmt.Println("======================== POSTING TUIT IN SQL CLOUD ========================")

    // Call BindJSON to bind the received JSON to
    if err := c.BindJSON(&newTuit); err != nil {
        return
    }

    err := sql.Init()

    if err!=nil{
        fmt.Println(err)
        c.JSON(http.StatusInternalServerError, "No connectado")
    }else{
        c.JSON(http.StatusCreated, "Conectado")
    }

}



func main() {

    fmt.Println("")
    fmt.Println(" ==========================  SERVIDOR  ========================== ")
    fmt.Println("")
	router := gin.Default()
    router.Use(gin.Recovery()) // Para recuperarse de Errores y enviar un 500

    router.POST("/addTuit/cloud/go", validateDataBases, postTuitCloud)

    router.Run()
}

//  /home/sopes1_s2_2021_g14/sopes1/SO1_G12_Proyecto1/Apis/GoApi/