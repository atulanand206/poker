package main

import (
	"fmt"
	"github.com/atulanand206/poker/src"
	"os"
	"log"
)

const dbFileName = "db.game.json"

func main() {
	store, closeStore, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer closeStore()

	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win.")
	game := poker.NewHoldem(store, poker.BlindAlerterFunc(poker.StdOutAlerter))
	poker.NewCli(os.Stdin, os.Stdout, game).PlayPoker()
}
