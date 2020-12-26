package poker_test

import (
	"github.com/atulanand206/poker/src"
	"testing"
	"strings"
	"bytes"
	"io"
)

var dummyStdOut = &bytes.Buffer{}

func TestCli(t *testing.T) {
	t.Run("it prints an error when a non numeric value is entered and game does not start", func(t *testing.T) {
		in := userSends("drama")
		out := &bytes.Buffer{}
		game := &poker.SpyGame{}

		cli := poker.NewCli(in, out, game)
		cli.PlayPoker()

		AssertGameNotStarted(t, game)
		AssertMessagesSentToUser(t, out, poker.PlayerPrompt, poker.BadPlayerInputErrMsg)
	})

	t.Run("starts the game with 7 players and finishes with Chris", func(t *testing.T) {
		out := &bytes.Buffer{}
		in := userSends("7", "Chris wins")
		game := &poker.SpyGame{}
		cli := poker.NewCli(in, out, game)
		cli.PlayPoker()

		AssertGameStartedWith(t, game, 7)
		AssertMessagesSentToUser(t, out, poker.PlayerPrompt)
		AssertGameFinishedWith(t, game, "Chris")
	})

	t.Run("starts the game with 3 players and finishes with Maria", func(t *testing.T) {
		out := &bytes.Buffer{}
		in := userSends("3", "Maria wins")
		game := &poker.SpyGame{}
		cli := poker.NewCli(in, out, game)
		cli.PlayPoker()

		AssertGameStartedWith(t, game, 3)
		AssertMessagesSentToUser(t, out, poker.PlayerPrompt)
		AssertGameFinishedWith(t, game, "Maria")
	})

	t.Run("it prints an error when the winner is declared incorrectly", func(t *testing.T) {
		out := &bytes.Buffer{}
		in := userSends("9", "you wish")
		game := &poker.SpyGame{}
		cli := poker.NewCli(in, out, game)
		cli.PlayPoker()

		AssertGameNotFinished(t, game)
		AssertMessagesSentToUser(t, out, poker.PlayerPrompt, poker.BadWinnerErrMsg)
	})
}

func userSends(messages ...string) io.Reader {
	return strings.NewReader(strings.Join(messages, "\n"))
}

func AssertMessagesSentToUser(t *testing.T, stdout *bytes.Buffer, messages ...string) {
	t.Helper()
	want := strings.Join(messages, "")
	got := stdout.String()
	if got != want {
		t.Errorf("got %q sent to stdout but expected %+v", got, messages)
	}
}

func AssertGameNotStarted(t *testing.T, game *poker.SpyGame) {
	t.Helper()
	if game.StartCalled {
		t.Errorf("game should not have started")
	}
}

func AssertGameStartedWith(t *testing.T, game *poker.SpyGame, players int) {
	t.Helper()
	if !game.StartCalled {
		t.Errorf("game should have started")
	}
	if game.StartedWith != players {
		t.Errorf("game should have started with %d players but started with %d", players, game.StartedWith)
	}
}

func AssertGameNotFinished(t *testing.T, game *poker.SpyGame) {
	t.Helper()
	if game.FinishedCalled {
		t.Errorf("game should not have finished")
	}
}

func AssertGameFinishedWith(t *testing.T, game *poker.SpyGame, winner string) {
	t.Helper()
	if game.FinishedWith != winner {
		t.Errorf("game should have finished with %s but got finished with %s", winner, game.FinishedWith)
	}
}
