package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("return fastest server", func(t *testing.T) {
		slowServer := makeTestServer(20 * time.Millisecond)
		fastServer := makeTestServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		want := fastUrl
		got, _ := Racer(slowUrl, fastUrl, 50*time.Millisecond)

		assertStrings(t, got, want)
	})

	t.Run("return an error by timeout", func(t *testing.T) {
		server := makeTestServer(20 * time.Millisecond)
		defer server.Close()

		_, err := Racer(server.URL, server.URL, 10*time.Millisecond)

		if err == nil {
			t.Errorf("expected an error")
		}
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("expected %q but got %q", want, got)
	}
}

func makeTestServer(delay time.Duration) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))

	return server
}
