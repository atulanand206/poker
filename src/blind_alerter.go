package poker

import (
	"time"
	"fmt"
	"os"
)

type Alert struct {
	ScheduledAt time.Duration
	Amount      int
}

type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int)
}

type BlindAlerterFunc func(duration time.Duration, amount int)

func (b BlindAlerterFunc) ScheduleAlertAt(duration time.Duration, amount int) {
	b(duration, amount)
}

func StdOutAlerter(duration time.Duration, amount int) {
	time.AfterFunc(duration, func() {
		fmt.Fprintf(os.Stdout, "Blind is now %d\n", amount)
	})
}

type SpyBlindAlerter struct {
	alerts []Alert
}

func (s *SpyBlindAlerter) Alerts() []Alert {
	return s.alerts
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.alerts = append(s.alerts, Alert{ScheduledAt: duration, Amount: amount})
}