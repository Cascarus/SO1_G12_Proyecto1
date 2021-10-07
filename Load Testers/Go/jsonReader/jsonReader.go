package jsonReader

import(
	"fmt"
	"encoding/json"
	"io/ioutil"
	ts "loadTester/types"
)

func ReadFile(path string) []ts.Tuit{

	f, err := ioutil.ReadFile(path)
	if err!=nil{
		fmt.Println(err)
	}
	
	var tuits[] ts.Tuit

	err = json.Unmarshal(f, &tuits)
	if err!=nil{
		fmt.Println(err)
	}

	return tuits
}