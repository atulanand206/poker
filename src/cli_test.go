package poker_test

import (
	"github.com/atulanand206/poker/src"
	"testing"
	"strings"
	"fmt"
)

func TestCli(t *testing.T) {
	cases := []struct {
		player string
	}{
		{"Natalie"},
		{"Rachel"},
	}

	for _, c := range cases {
		t.Run(recordTestName(c.player), func(t *testing.T) {
			store := &poker.StubPlayerStore{}
			in := strings.NewReader(fmt.Sprintf("%s wins\n", c.player))
			cli := poker.NewCli(store, in)
			cli.PlayPoker()
			poker.AssertPlayerWin(t, store, c.player)
		})
	}
}

func recordTestName(player string) string {
	return fmt.Sprintf("record %s win from user input", player)
}
