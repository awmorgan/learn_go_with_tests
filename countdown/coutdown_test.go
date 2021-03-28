package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	b := &bytes.Buffer{}
	Countdown(b)
	got := b.String()
	want := "3"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}