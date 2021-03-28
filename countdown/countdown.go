package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type DefaultSleeper struct {
}

func (s *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func main() {
	Countdown(os.Stdout, &DefaultSleeper{})
}

func Countdown(w io.Writer, s Sleeper) {
	const countdownStart = 3
	for i := countdownStart; i > 0; i-- {
		s.Sleep()
		fmt.Fprintf(w, "%d\n", i)
	}
	const finalWord = "Go!"
	s.Sleep()
	fmt.Fprintf(w, "%s\n", finalWord)
}

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (s *ConfigurableSleeper) Sleep()  {
	s.sleep(s.duration)
}
