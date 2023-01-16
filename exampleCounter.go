package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"io/ioutil"
	"net/http"
	"sync"
)

var exMutex = &sync.Mutex{}

type Handler interface {
	webCounter(http.ResponseWriter, *http.Request)
}

func readFile() []byte {
	jsonResult, _ := ioutil.ReadFile("person.json")
	return jsonResult
}
func webCounter(w http.ResponseWriter, r *http.Request) {
	exMutex.Lock()
	if r.Method == "POST" {
		var s1 dataPerson
		s2, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(s2, &s1)
		validate := validator.New()
		err := validate.Struct(s1)
		validatorErr := err.(validator.ValidationErrors)
		if validatorErr == nil {

			fmt.Fprintf(w, string(makePerson(s1)))
		}

		exMutex.Unlock()
	}
}
