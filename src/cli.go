package poker

import (
	"io"
	"bufio"
	"strings"
	"fmt"
	"strconv"
	"errors"
)

type Cli struct {
	input  *bufio.Scanner
	output io.Writer
	game   Game
}

func NewCli(input io.Reader, output io.Writer, game Game) *Cli {
	return &Cli{bufio.NewScanner(input), output, game}
}

const PlayerPrompt = "Please enter the number of players: "
const BadPlayerInputErrMsg = "Bad value received for number of players, please try again with a number"
const BadWinnerErrMsg = "Bad value received for winner, please enter in the specified format"

func (c *Cli) PlayPoker() {
	fmt.Fprint(c.output, PlayerPrompt)
	numberOfPlayers, err := strconv.Atoi(c.readLine())
	if err != nil {
		fmt.Fprint(c.output, BadPlayerInputErrMsg)
		return
	}
	c.game.Start(numberOfPlayers)
	winner, err := extractWinner(c.readLine())
	if err != nil {
		fmt.Fprint(c.output, BadWinnerErrMsg)
		return
	}
	c.game.Finish(winner)
}

func extractWinner(userInput string) (string, error) {
	if !strings.Contains(userInput, "wins") {
		return "", errors.New(BadWinnerErrMsg)
	}
	return strings.Replace(userInput, " wins", "", 1), nil
}

func (cli *Cli) readLine() string {
	cli.input.Scan()
	return cli.input.Text()
}
