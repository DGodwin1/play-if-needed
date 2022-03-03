package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/sort-pairings", Doubles)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Doubles(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		//You should only be able to post to the backend.
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	//Let's print out the players.
	for _, player := range strings.Split(r.FormValue("players"), " ") {
		fmt.Println(player)
	}
	w.WriteHeader(200)
}
