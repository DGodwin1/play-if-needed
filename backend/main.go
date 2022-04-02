package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/sort-pairings", Doubles)
	log.Fatal(http.ListenAndServe(":9001", nil))
}

func Doubles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	w.Header().Set("Content-Type", "application/json")

	json, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	fmt.Println(string(json))

}
