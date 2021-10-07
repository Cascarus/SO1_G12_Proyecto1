package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"time"
	"sync"
	"net/http"
	"os"
	"log"
	"bytes"
	jr "loadTester/jsonReader"
	ts "loadTester/types"
)


func loadTuits(jobs <-chan ts.Tuit, results chan<- string, db string){
	
	for n:= range jobs {

		if db=="azure" {
			postNewTuitCosmos(n)
			
		}else{
			postNewTuitCloud(n)
		}
		
		results <- n.Nombre+" Loaded"
	}
}


func postNewTuitCosmos(tuit ts.Tuit){

	url := os.Getenv("API_HOST")+"/addTuit/cosmos/go"
    fmt.Println("URL:>", url)

    data, err := json.Marshal(tuit)
	if err != nil{
		log.Fatal(err)
	}

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Print(err)
    }
    defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var newLog ts.Log
	json.Unmarshal([]byte(string(body)), &newLog)
	logs = append(logs, newLog)

	/*if resp.StatusCode != 500 {
		fmt.Println("Tuit loaded:", string(body))
	}else{
		fmt.Println("Error loading tuit: ", tuit)
	}*/
	
	return
}


func postNewTuitCloud(tuit ts.Tuit){

	url := os.Getenv("API_HOST")+"/addTuit/cloud/go"
    fmt.Println("URL:>", url)

    data, err := json.Marshal(tuit)
	if err != nil{
		log.Fatal(err)
	}

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Print(err)
    }
    defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var newLog ts.Log
	json.Unmarshal([]byte(string(body)), &newLog)
	logs = append(logs, newLog)

	/*if resp.StatusCode != 500 {
		fmt.Println("Tuit loaded:", string(body))
	}else{
		fmt.Println("Error loading tuit: ", tuit)
	}*/
	
	return
}

func getData(){

	resp, err := http.Get(os.Getenv("API_HOST")+"/getTuits/go")

	if err!= nil{
		log.Fatal(err)
	}
	
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(body))
}



func start(){

	resp, err := http.Get(os.Getenv("API_HOST")+"/startLoad/go")

	if err!= nil{
		log.Fatal(err)
	}
	
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        log.Fatal(err)
    }
	fmt.Println(string(body))
}


func finish(){

	resp, err := http.Get(os.Getenv("API_HOST")+"/closeLoad/go")

	if err!= nil{
		log.Fatal(err)
	}
	
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        log.Fatal(err)
    }
	fmt.Println(string(body))
}

var logs[] ts.Log
var tuits[] ts.Tuit
func main() {
	
	fmt.Println("==================================== STARTING ====================================")
	fmt.Println("")
	//ps.InitPubSub()

//======================================== READING JSON FILE

	fmt.Println("File will be read in 5 seconds...")
	time.Sleep(5 * time.Second)
	uno()
	


//======================================== LOADING TUITS

	fmt.Println("")
	start()
	fmt.Println("Tuits are going to be loaded in 5 seconds...")
	time.Sleep(5 * time.Second)

	/*newTuit := ts.Tuit{
		Nombre: "Efrain Alvarez",
		Comentario: "Cualquier cosa de preuba",
		Fecha: "30/7/2021",
		Hashtags: []string{"coso1","coso2"},
		Upvotes: 110,
		Downvotes: 66,
	}
	fmt.Println(newTuit)*/

	dos()
	finish()
	fmt.Println(logs)
	//fmt.Println(tuits)
	//loadTuits(tuits[0])
}

func uno(){

	wg:= &sync.WaitGroup{}
	wg.Add(1)
	
	go func(){
		tuits= jr.ReadFile("./prueba.json")
		fmt.Println(tuits)
		wg.Done()
	}()
	wg.Wait()

}


func dos(){

	jobs := make(chan ts.Tuit, len(tuits))
	results := make(chan string, len(tuits))

	go loadTuits(jobs, results, "azure")
	go loadTuits(jobs, results, "google")

	for i := 0; i < len(tuits); i++ {
		jobs <- tuits[i]
	}
	close(jobs)

	for j := 0; j < len(tuits); j++ {
		fmt.Println(<-results)
	}

}