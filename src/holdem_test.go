package poker_test

import (
	"github.com/atulanand206/poker/src"
	"testing"
	"time"
	"fmt"
)

func TestHoldem_Start(t *testing.T) {
	t.Run("it schedules printing of blind values", func(t *testing.T) {
		playerStore := &poker.StubPlayerStore{}
		blindAlerter := &poker.SpyBlindAlerter{}
		game := poker.NewHoldem(playerStore, blindAlerter)
		game.Start(5)

		cases := []poker.Alert{
			{0 * time.Second, 100},
			{10 * time.Second, 200},
			{20 * time.Second, 300},
			{30 * time.Second, 400},
			{40 * time.Second, 500},
			{50 * time.Second, 600},
			{60 * time.Second, 800},
			{70 * time.Second, 1000},
			{80 * time.Second, 2000},
			{90 * time.Second, 4000},
			{100 * time.Second, 8000},
		}

		AssertScheduling(t, cases, blindAlerter)
	})

	t.Run("it schedules printing of blind values with more players", func(t *testing.T) {
		playerStore := &poker.StubPlayerStore{}
		blindAlerter := &poker.SpyBlindAlerter{}
		game := poker.NewHoldem(playerStore, blindAlerter)
		game.Start(7)

		cases := []poker.Alert{
			{0 * time.Second, 100},
			{12 * time.Second, 200},
			{24 * time.Second, 300},
			{36 * time.Second, 400},
		}

		AssertScheduling(t, cases, blindAlerter)
	})
}

func AssertScheduling(t *testing.T, cases []poker.Alert, blindAlerter *poker.SpyBlindAlerter) {
	for i, c := range cases {
		t.Run(scheduleTestName(c), func(t *testing.T) {
			alerts := blindAlerter.Alerts()
			if len(alerts) <= i {
				t.Fatal(fmt.Sprintf("alert %d was not scheduled %v", c.Amount, c.ScheduledAt))
			}
			alert := alerts[i]
			AssertScheduleAlert(t, alert, c)
		})
	}
}

func scheduleTestName(i poker.Alert) string {
	return fmt.Sprintf("%d expected at %v", i.ScheduledAt, i.Amount)
}

func AssertScheduleAlert(t *testing.T, alert poker.Alert, c poker.Alert) {
	t.Helper()
	gotAmount := alert.Amount
	if gotAmount != c.Amount {
		t.Errorf("got amount %d, want %d", gotAmount, c.Amount)
	}
	gotScheduleTime := alert.ScheduledAt
	if gotScheduleTime != c.ScheduledAt {
		t.Errorf("got scheduled time of %v, want %v", gotScheduleTime, c.ScheduledAt)
	}
}

func TestHoldem_Finish(t *testing.T) {
	t.Run("it finishes the game with a user", func(t *testing.T) {
		cases := []struct {
			player string
		}{
			{"Natalie"},
			{"Rachel"},
		}

		for _, c := range cases {
			t.Run(recordTestName(c.player), func(t *testing.T) {
				playerStore := &poker.StubPlayerStore{}
				blindAlerter := &poker.SpyBlindAlerter{}
				game := poker.NewHoldem(playerStore, blindAlerter)
				game.Finish(c.player)
				poker.AssertPlayerWin(t, playerStore, c.player)
			})
		}
	})
}

func recordTestName(player string) string {
	return fmt.Sprintf("record %s win from user input", player)
}
