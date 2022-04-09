package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
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

	//for _, v := range players {
	//	//fmt.Println(v.Name, v.Rank)
	//}

	remaining := len(players) % 4
	if remaining > 1 {
		//You have a difficult number of players to deal with. So let's handle it.
		players, leftovers := makeThePairingsEasier(players, remaining)
		fmt.Println(players)
		fmt.Println(leftovers)
		//do what we need to do with leftovers / players here. or at end?
	}
	//We have a nice number of players. Let's get to shuffling.

	//Just a basic shuffle.
	fmt.Printf("Got %v\n", players)
	fmt.Println("TIME TO SHUFFLE")
	rand.Shuffle(len(players), func(i, j int) { players[i], players[j] = players[j], players[i] })
	fmt.Println(players)

}

func makeThePairingsEasier(players []Player, remaining int) ([]Player, []Player) {
	//Based on a difficult number of players, will return a modified list of players in a normal doubles
	//format and any 'leftovers' in a second Player array.
	switch remaining {
	case 1:
		fmt.Println("Someone has to sit out. Sorry. Going to randomly pick someone.")
		player := players[rand.Intn(len(players))]
		fmt.Printf("%v is going to have to sit out.", player.Name)
	case 2:
		fmt.Println("A singles can take place. Picking two people.")
	case 3:
		fmt.Println("We need one more for another doubles or someone can sit out. 3 people to chat.")
	}
	return []Player{}, []Player{}
}
