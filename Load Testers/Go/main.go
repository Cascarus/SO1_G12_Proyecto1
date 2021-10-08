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

	t := time.Now() // EMPIEZA EL TIEMPO
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

	if resp.StatusCode != 500 {
		cosmosLogs = append(cosmosLogs, ts.Log{ StatusNumber:http.StatusCreated, Message:string(body), Time:time.Since(t) } )
		fmt.Println("Tuit loaded:", string(body))
	}else{
		fmt.Println("Error loading tuit: ", tuit)
	}
	
	return
}


func postNewTuitCloud(tuit ts.Tuit){

	url := os.Getenv("API_HOST")+"/addTuit/cloud/go"
    fmt.Println("URL:>", url)

    data, err := json.Marshal(tuit)
	if err != nil{
		log.Fatal(err)
	}

	t := time.Now() // EMPIEZA EL TIEMPO
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

	if resp.StatusCode != 500 {
		cloudLogs = append(cloudLogs, ts.Log{ StatusNumber:http.StatusCreated, Message:string(body), Time:time.Since(t) } )
	}else{
		fmt.Println("Error loading tuit: ", tuit)
	}
	
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


func finish(msg ts.Message){

	json_data, err := json.Marshal(msg)
    if err != nil {
        return
    }

	resp, err := http.Post(os.Getenv("API_HOST")+"/closeLoad/go", "application/json", bytes.NewBuffer(json_data))

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


func sendResults(){

	var timeCosmos time.Duration
    var timeCloud time.Duration

    for i := 0; i < len(cosmosLogs); i++ {
        timeCosmos+=cosmosLogs[i].Time
    }

    finish(ts.Message{ Guardados:len(cosmosLogs), Api:"Go", TiempoCarga: fmt.Sprint(timeCosmos), Db:"Azure Cosmos"})


    for i := 0; i < len(cloudLogs); i++ {
        timeCloud+=cloudLogs[i].Time
    }
    finish(ts.Message{ Guardados:len(cloudLogs), Api:"Go", TiempoCarga: fmt.Sprint(timeCloud), Db:"Cloud SQL"})

}


var tuits[] ts.Tuit

var cosmosLogs[] ts.Log
var cloudLogs[] ts.Log
var logs[] ts.Log


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


	dos()
	sendResults()
	fmt.Println(cosmosLogs)
	fmt.Println(cloudLogs)
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
	//go loadTuits(jobs, results, "google")

	for i := 0; i < len(tuits); i++ {
		jobs <- tuits[i]
	}
	close(jobs)

	for j := 0; j < len(tuits); j++ {
		fmt.Println(<-results)
	}

}