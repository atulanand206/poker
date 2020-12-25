package main

import (
	"github.com/atulanand206/poker/cmd"
	"net/http"
	"os"
)

const dbFileName = "db.game.json"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_CREATE|os.O_RDWR, 0666)
	poker.ErrorFileOpening(err, dbFileName)
	store, err := poker.NewFileSystemPlayerStore(db)
	poker.ErrorFileCreation(err)
	server := poker.NewPlayerServer(store)
	handler := http.HandlerFunc(server.ServeHTTP)
	err = http.ListenAndServe(":5000", handler)
	poker.ErrorListenAndServe(err)
}