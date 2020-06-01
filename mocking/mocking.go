package mocking

import (
	"fmt"
	"io"
	"time"
)

const (
	finalWord      = "Go!"
	countdownStart = 3
	sleep          = "sleep"
	write          = "write"
)

//Sleeper
type Sleeper interface {
	Sleep()
}

//SpySleeper
type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

//CountdownOperationsSpy
type CountdownOperationsSpy struct {
	Calls []string
}

func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

//ConfigurableSleeper
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

// SpyTime
type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

// CountDown function
func CountDown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
	}

	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
	}

	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}
