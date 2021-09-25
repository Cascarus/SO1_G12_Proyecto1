package pubSub

import (
	"fmt"
	"os"
	"errors"
	"log"
	"context"
	"cloud.google.com/go/pubsub"
	ts "goApi/types"
	"strconv"
)

var topic *pubsub.Topic

func InitPubSub(message ts.Message) {

	err:= start()

	if err != nil{
		fmt.Println("Not connected to topic")
	}else{
		fmt.Println("Connected to topic")
		err:= publishMessage(message)
		topic.Stop()

		if err!= nil{
			fmt.Println("Message not published")
		}
	}
}



func start() error {

	fmt.Println("================================ STARTING")

	ctx := context.Background()

	client, err := pubsub.NewClient(ctx, os.Getenv("PROYECT"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	topicName := "olympics"
	topic = client.Topic(topicName)

	// Create the topic if it doesn't exist.
	exists, err := topic.Exists(ctx)
	if err != nil {
		log.Print(err)
		return nil
	}
	if !exists {
		log.Printf("Topic %v doesn't exist - creating it", topicName)
		_, err = client.CreateTopic(ctx, topicName)
		if err != nil {
			return errors.New("Topic couldn't be created...")
		}
	}

	return nil
}


func publishMessage(message ts.Message) error {

	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, os.Getenv("PROYECT"))
	if err != nil {
		fmt.Println("Error 1")
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}
	defer client.Close()

	t := client.Topic("olympics")
	result := t.Publish(ctx, &pubsub.Message{
			Data: []byte("Load completed!"),
			Attributes: map[string]string{
					"guardados":   strconv.Itoa(message.Guardados),
					"api": message.Api,
					"tiempoCarga":   message.TiempoCarga,
					"db":   message.Db,
			},
	})
	/*
	Guardados int `json:"guardados"`
	api string `json:"api"`
	tiempoCarga string `json:"tiempoCarga"`
	db string `json:"db"`
	*/

	// Block until the result is returned and a server-generated
	// ID is returned for the published message.
	id, err := result.Get(ctx)
	if err != nil {
		fmt.Println("Error 2")
		log.Print(err)
		return fmt.Errorf("Get: %v", err)
	}
	fmt.Println("Published message with custom attributes; msg ID: %v\n", id)
	return nil

}
