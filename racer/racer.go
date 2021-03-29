package racer

import (
	"net/http"
	"time"
)

func Racer(a, b string) string {
	if timeHttpGet(a) < timeHttpGet(b) {
		return a
	}
	return b
}

func timeHttpGet(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	duration := time.Since(start)
	return duration
}
