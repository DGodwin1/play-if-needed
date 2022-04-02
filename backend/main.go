package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/sort-pairings", Doubles)
	log.Fatal(http.ListenAndServe(":9001", nil))
}

type Player struct {
	Name string `json:"name"`
	Rank int    `json:"rank"`
}

func Doubles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	w.Header().Set("Content-Type", "application/json")

	j, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	var players []Player

	err = json.Unmarshal(j, &players)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range players {
		fmt.Println(v.Name, v.Rank)
	}

}
