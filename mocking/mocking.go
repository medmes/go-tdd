package mocking

import (
	"fmt"
	"io"
	"time"
)

const (
	finalWord      = "Go!"
	countdownStart = 3
)

func CountDown(out io.Writer) {
	for i := countdownStart; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintln(out, i)
	}
	time.Sleep(1 * time.Second)
	fmt.Fprint(out, finalWord)
}
