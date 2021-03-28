package main

import (
	"bytes"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCountdown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
		b := &bytes.Buffer{}
		s := &CountdownOperationsSpy{}
		Countdown(b, s)
		got := b.String()
		want := `3
2
1
Go!
`
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("sleep before every print", func(t *testing.T) {
		s := &CountdownOperationsSpy{}
		Countdown(s, s)
		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !cmp.Equal(want, s.Calls) {
			t.Errorf("wanted calls %v got %v", want, s.Calls)
		}
	})
}

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

const write = "write"
const sleep = "sleep"
