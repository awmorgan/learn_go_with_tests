package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("slow, fast", func(t *testing.T) {
		slowServer := server(5 * time.Millisecond)
		fastServer := server(1 * time.Millisecond)
		defer slowServer.Close()
		defer fastServer.Close()

		want := fastServer.URL
		got, err := Racer(slowServer.URL, fastServer.URL)
		assertRacer(got, want, err, t)
	})
	t.Run("fast, slow", func(t *testing.T) {
		slowServer := server(5 * time.Millisecond)
		fastServer := server(1 * time.Millisecond)
		defer slowServer.Close()
		defer fastServer.Close()

		want := fastServer.URL
		got, err := Racer(fastServer.URL, slowServer.URL)
		assertRacer(got, want, err, t)
	})
	t.Run("error after 10s", func(t *testing.T) {
		timeout := 10 * time.Millisecond
		slowServer := server(timeout + 1*time.Millisecond)
		fastServer := server(timeout + 1*time.Nanosecond)
		defer slowServer.Close()
		defer fastServer.Close()

		_, err := RacerWithTimeout(fastServer.URL, slowServer.URL, timeout)
		if err == nil {
			t.Error("expected error but didn't get one")
		}
	})

}

func assertRacer(got string, want string, err error, t testing.TB) {
	t.Helper()
	if err != nil {
		t.Errorf("did not expect error: %v", err)
	}
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
