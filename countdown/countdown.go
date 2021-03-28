package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	Countdown(os.Stdout, &ConfigurableSleeper{duration: 1 * time.Second, sleep: time.Sleep})
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

func (s *ConfigurableSleeper) Sleep() {
	s.sleep(s.duration)
}
