package main

import (
	//"io/ioutil"
	"net/http"
	"fmt"
	"errors"
	"log"
	"context"
	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
	"golang.org/x/oauth2/google"
)

var topic *pubsub.Topic


func InitPubSub(w http.ResponseWriter, r *http.Request) {

    fmt.Println("======================== SENDING MESSAGE ========================")

	err := publishMessage()

	if err != nil{
		fmt.Fprint(w, "Not connected to topic")
		return
	}else{
		fmt.Println(w, "Connected to topic")
		err:= publishMessage()
		topic.Stop()

		if err!= nil{
			fmt.Println(w, "Message not published")
		}
		return
	}


}

func publishMessage() error {

	ctx := context.Background()

	/*data, er := ioutil.ReadFile("PS.json")
	if er != nil {
		log.Fatal(er)
	}*/

	data := `{
		"type": "service_account",
		"project_id": "deft-idiom-324423",
		"private_key_id": "ee70ee13d97c063d13b038926488cbce3da8c1fe",
		"private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDSRvkRwg2rvoGg\nTc0HleEIVxoLdZ8HKxw7LJwVn97x8kIrpJQWtFig/njLQyd5QVQ65eqKbnAQWyG4\nZI3ogiC9M7gFX05i5zHND2dH5JXwvviS5pTYa0/aM6Yb3MbQi88ZGnRhbeowXgme\njuywXF5QLi3fn2QpQ3IoN+kDuZ3W8CxcAtGcLKgchaEO5MUnGxT2YxZPUCoXodQh\nzSQBDvpL8BrscxP3fBIFbqwWt/cDnJWd28qjukZs5E+L/VXetNlnJo/4dsfBAo6g\nZ3KZISIBOSd4aAcgmfidd2oUF3e7k12RMV8TwYbzyC0Gj7yhhMYCXI0/J0dKwQEP\nf5qorTyRAgMBAAECggEAGLHGEVxRP7ClR36ri+8BmPmlsN/U18Ir1BU2lWGLjK1v\nMPHohEwUtn98Dx7pVVejPTGAHFbA6WLScHW6pqoVjzKyM0tQiNPu6M6cyfgh1b6P\nsazLoSjlHVKAePgyDw7EOQX+0exXGuwIRCszx7hpSRBLSd5NpHzrvKZoQow7aaDf\nhjxS3eRGiHkCMcsJV9fD6SZLocpdTmuwcvqrTGY1Sxn0O68OMG6IEDVNy/mxXWWL\nOoVRdNBjQo3MJqFlpz4TJ4+jjFoIv5IYbUZqdGj30yoTisEDunvmbYVpliobCVSC\nKdhNua8mdCrCLElipOFiHcNYXhOczQtKyWfb+vPpswKBgQDv3Vunn5RcULOfK4+f\nuguTMkeDybmy5/wWOEiHKirv/3HdW1+2nZE2VaMNcnHfzWH7kIVXoAW9nCzOMZ9j\nqmX794aiqqgJduR2Pw2Mqi4fOhFJUOBa2dOvVHeCESWbNnxWswGbM3tCtBWHmie7\nEoKbjrV8+yWe+R7CUO1dBjxodwKBgQDgbBkQ3ef19KXdAde6T2dtHP5CX95w6CJx\n5zgRXr+pdpPqhhszuXRt7mHMAGjWh/kWnZwIq/+0LJFMFiUxIZL0mTE6UnYloXOp\nVv7NiqaRN68WShqTGObzzgTom497lI4jR0dP7POObcqu5s7apmhDZK+fxs84TefE\nnet8dLFNNwKBgQC6t83JRmfvFMM+fGJpLCImi1UwOa/cnMmXYmjTDvgtquOwNJjl\nRvLrIO60YQpT9UT41x91fpP5bcFTIT26D8MjySN9LOtxsqNViO+7OB5/IGykbdi3\n4Cjwqwf8r+xeTqOrudzeO80Pt6+qx012SopxHT4Z9Ebs4XAYQ3cCmwAbMQKBgGuD\nvhpzhR4y/4c6y8QJKG6AtlrMHQAQZfgVoqnHr6CbG0/+wWdtUJcd3iJii9dDOxUX\nmtoYtJ73vwApl9XK1OFzxr6/JLTwfT3CXL3Rz+zANZRDGioggvyIVZeudvXofJPw\nIPzsct5oQPK7xpu/nzGyOeUc1MePoxpx4ZA1Q3/PAoGAJoDIL53k8NNBAexSfCnq\ngHYcgiVSghtXZ/IuSkxLjBEtkuTCGC1tOEbWkkuq8Jvk9RI6zPozlCJFuuVSfyMu\nF/Lj0oQPTd5AxjKS4YyqqrHEVnSX+1pQID/52flxp1rs0wXzHakH/H9lAc/w000v\nnBDjctxGNFqwr08IJ9P9r9g=\n-----END PRIVATE KEY-----\n",
		"client_email": "pubsub-system@deft-idiom-324423.iam.gserviceaccount.com",
		"client_id": "110651777843878913217",
		"auth_uri": "https://accounts.google.com/o/oauth2/auth",
		"token_uri": "https://oauth2.googleapis.com/token",
		"auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
		"client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/pubsub-system%40deft-idiom-324423.iam.gserviceaccount.com"
	  }`

	creds, err := google.CredentialsFromJSON(ctx, []byte(data), pubsub.ScopePubSub)
	if err != nil {
		return errors.New("Nuevo error")
		// TODO: handle error.
	}
	client, err := pubsub.NewClient(ctx, "deft-idiom-324423", option.WithCredentials(creds))

	if err != nil {
		fmt.Println("Error 1")
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}
	defer client.Close()

	t := client.Topic("olympics")
	result := t.Publish(ctx, &pubsub.Message{
			Data: []byte("Load completed!"),
			Attributes: map[string]string{
					"guardados":   "strconv.Itoa(message.Guardados)",
					"api": "message.Api",
					"tiempoCarga":   "message.TiempoCarga",
					"db":   "message.Db",
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


func main(){

	fmt.Println("Hola Mundo!")
	
	publishMessage()

}
