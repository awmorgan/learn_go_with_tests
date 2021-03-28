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
	fmt.Fprintf(w, "3")
}
