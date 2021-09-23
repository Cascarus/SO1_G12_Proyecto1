package pubSub

import (
	"fmt"
	"os"
	"errors"

	"log"
	"context"
	"cloud.google.com/go/pubsub"
)

var topic *pubsub.Topic

func InitPubSub() {

	err:= start()

	if err != nil{
		fmt.Println("Not connected to topic")
	}else{
		fmt.Println("Connected to topic")
		err:= publishMessage(os.Getenv("PROYECT"), "olympics")
		//topic.Stop()

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


func publishMessage(projectID, topicID string) error {

	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		fmt.Println("Error 1")
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}
	defer client.Close()

	t := client.Topic(topicID)
	result := t.Publish(ctx, &pubsub.Message{
			Data: []byte("Hello world!"),
			Attributes: map[string]string{
					"origin":   "golang",
					"username": "gcp",
			},
	})
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
