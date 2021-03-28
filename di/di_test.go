package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Bill")

	got := buffer.String()
	want := "Hello, Bill"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
