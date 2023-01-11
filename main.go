package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var mutex = &sync.Mutex{}
var counter int

func inCounter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock() //зачем нам нужен мутекс?
	counter++
	fmt.Fprintf(w, strconv.Itoa(counter))
	fmt.Println(counter)
	mutex.Unlock()
}

func main() {
	http.HandleFunc("/", inCounter)
	log.Fatal(http.ListenAndServe(":8083", nil))
}
