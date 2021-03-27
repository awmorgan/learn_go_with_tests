package main

import "testing"

func TestHello(t *testing.T)  {
	got := Hello("Bill")
	want := "Hello, Bill"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}