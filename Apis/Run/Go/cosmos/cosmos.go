package cosmos

import (
	"fmt"
	"errors"
	"os"
	"time"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	ts "goApi/types"
)


// connects to MongoDB
func Connect() (*mongo.Client, error) {
	
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_CS")).SetDirect(true)
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
func Create(tuit ts.Tuit) (string, error) {

	c, err := Connect()
    
    if err!=nil {
        fmt.Println(err)
        return "", err
    }

    ctx := context.Background()
    defer c.Disconnect(ctx)

    todoCollection := c.Database(os.Getenv("DB")).Collection(os.Getenv("COLLECTION"))
    r, err := todoCollection.InsertOne(ctx, tuit)
    if err != nil {
        return "", errors.New("Failed to load tuit")
    }
    //fmt.Println("added todo", r.InsertedID)

	return "Tuit loaded successfully with id: "+fmt.Sprint(r.InsertedID), nil
}