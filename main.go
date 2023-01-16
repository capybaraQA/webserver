package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type dataPerson struct {
	Name string `validate:"required"`
	Age  int    `validate:"required"`
}

var data = make([]dataPerson, 0)

func makePerson(sd dataPerson) []byte {
	/*if sd.Name == "" || sd.Age == 0 {
		return nil
	}*/
	data = append(data, sd)
	createJson, _ := json.MarshalIndent(data, "", " ")
	_ = ioutil.WriteFile("person.json", createJson, 0644)
	return createJson
}

func main() {

	http.HandleFunc("/add", webCounter)
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		exMutex.Lock()
		if request.Method == "GET" {
			fmt.Fprintf(writer, string(readFile()))
		}
		exMutex.Unlock() //?

	})

	log.Fatal(http.ListenAndServe(":8083", nil))

}
