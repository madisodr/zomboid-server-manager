package utils

import "time"

type SecondsTimer struct {
	Timer *time.Timer
	End   time.Time
}

func NewSecondsTimer(t time.Duration) *SecondsTimer {
	return &SecondsTimer{time.NewTimer(t), time.Now().Add(t)}
}

func (s *SecondsTimer) Reset(t time.Duration) {
	s.Timer.Reset(t)
	s.End = time.Now().Add(t)
}

func (s *SecondsTimer) Stop() {
	s.Timer.Stop()
}

func (s *SecondsTimer) TimeRemaining() time.Duration {
	return s.End.Sub(time.Now())
}
