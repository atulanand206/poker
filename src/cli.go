package poker

import (
	"io"
	"bufio"
	"strings"
)

type Cli struct {
	store PlayerStore
	input *bufio.Scanner
}

func NewCli(store PlayerStore, input io.Reader) *Cli {
	return &Cli{store, bufio.NewScanner(input)}
}

func (c *Cli) PlayPoker() {
	c.input.Scan()
	c.store.RecordWin(extractWinner(c.input.Text()))
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}
