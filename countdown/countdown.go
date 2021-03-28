package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	Countdown(os.Stdout)
}

func Countdown(w io.Writer) {
	const countdownStart = 3
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintf(w, "%d\n", i)
	}
	const finalWord = "Go!"
	fmt.Fprintf(w, "%s\n", finalWord)
}
