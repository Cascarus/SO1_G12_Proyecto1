package cosmos

import (
	"fmt"
	"errors"
	"net/http"
	"time"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type Tuit struct {

    ID primitive.ObjectID `json:"_id,omitempty"`
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


func PostTuitCosmos(w http.ResponseWriter, r *http.Request) {
    t := time.Now()
    var newTuit Tuit

    fmt.Println("======================== POSTING TUIT IN COSMOS ========================")


	if err := json.NewDecoder(r.Body).Decode(&newTuit); err != nil {
		fmt.Fprint(w, "Error reading Tuit!")
		return
	}
	msg, err := Create(newTuit)
	if err!=nil{
        fmt.Println(err)
		fmt.Fprint(w, Log{StatusNumber:http.StatusInternalServerError, Message:fmt.Sprint(err), Time:time.Since(t) })
    }else{
		fmt.Fprint(w, Log{ StatusNumber:http.StatusCreated, Message:msg, Time:time.Since(t) })
    }

}



// connects to MongoDB
func Connect() (*mongo.Client, error) {
	
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb://go-mongo:NWm0Ub0V1DZxVOLZb6IyMPa4a6HCHqWEuj8DZhjHV9VVFScnSWFDk0ky2xX61sZemUeq7Q61Tv4stiJKYVrXNw==@go-mongo.mongo.cosmos.azure.com:10255/?ssl=true&replicaSet=globaldb&retrywrites=false&maxIdleTimeMS=120000&appName=@go-mongo@").SetDirect(true)
	c, err := mongo.NewClient(clientOptions)

	err = c.Connect(ctx)

	if err != nil {
        return nil, errors.New("Unable to initialize database connection")
		//log.Fatalf("unable to initialize connection %v", err)
	}
	err = c.Ping(ctx, nil)
	if err != nil {
        return nil, errors.New("Unable to connect to database")
		//log.Fatalf("unable to connect %v", err)
	}

	return c, nil
}



// creates a todo
func Create(tuit Tuit) (string, error) {

	c, err := Connect()
    
    if err!=nil {
        fmt.Println(err)
        return "", err
    }

    ctx := context.Background()
    defer c.Disconnect(ctx)

    todoCollection := c.Database("Olympics").Collection("Tuits")
    r, err := todoCollection.InsertOne(ctx, tuit)
    if err != nil {
        return "", errors.New("Failed to load tuit")
    }
    //fmt.Println("added todo", r.InsertedID)

	return "Tuit loaded successfully with id: "+fmt.Sprint(r.InsertedID), nil
}