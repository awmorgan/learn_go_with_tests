package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	slowServer := server(5*time.Millisecond)
	fastServer := server(1*time.Millisecond)

	want := fastServer.URL
	got := Racer(slowServer.URL, fastServer.URL)
	assertRacer(got, want, t)
	
	want = fastServer.URL
	got = Racer(fastServer.URL, slowServer.URL)
	assertRacer(got, want, t)

	slowServer.Close()
	fastServer.Close()
}

func assertRacer(got string, want string, t testing.TB) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func server(sleep time.Duration) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if sleep.Nanoseconds() > 0 {
			time.Sleep(sleep)
		}
		rw.WriteHeader(http.StatusOK)
	}))
	return server
}
