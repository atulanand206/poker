package main

import (
	"github.com/atulanand206/poker/src"
	"net/http"
	"log"
)

const dbFileName = "db.game.json"

func main() {
	store, closeStore, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer closeStore()

	server, _ := poker.NewPlayerServer(store)
	handler := http.HandlerFunc(server.ServeHTTP)
	err = http.ListenAndServe(":5000", handler)
	err = poker.ErrorListenAndServe(err)
}