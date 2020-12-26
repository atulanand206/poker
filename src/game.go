package poker

import "time"

type Game interface {
	Start(players int)
	Finish(input string)
}

type Holdem struct {
	store   PlayerStore
	alerter BlindAlerter
}

func NewHoldem(store PlayerStore, alerter BlindAlerter) *Holdem {
	return &Holdem{store, alerter}
}

func (g *Holdem) Start(players int) {
	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	blindIncrement := time.Duration(5+players) * time.Second
	for _, blind := range blinds {
		g.alerter.ScheduleAlertAt(blindTime, blind)
		blindTime += blindIncrement
	}
}

func (g *Holdem) Finish(input string) {
	g.store.RecordWin(input)
}

type SpyGame struct {
	StartCalled bool
	StartedWith int

	FinishedCalled bool
	FinishedWith   string
}

func (s *SpyGame) Start(players int) {
	s.StartedWith = players
	s.StartCalled = true
}

func (s *SpyGame) Finish(input string) {
	s.FinishedWith = input
	s.FinishedCalled = true
}
