package poker

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"fmt"
)

func TestGetPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Angelina": 20,
			"Floyd":    10,
		},
		nil, nil,
	}
	cases := []struct {
		player string
		score  int
	}{
		{"Angelina", 20},
		{"Floyd", 10},
	}
	server, _ := NewPlayerServer(&store)
	for _, c := range cases {
		t.Run(testName(c.player), func(t *testing.T) {
			request := getRequest(c.player)
			response := httptest.NewRecorder()
			server.ServeHTTP(response, request)
			AssertEqual(t, response.Body.String(), fmt.Sprint(c.score))
			AssertEqualIntegers(t, response.Code, http.StatusOK)
		})
	}
	t.Run("404 on missing player", func(t *testing.T) {
		request := getRequest("Apollo")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		AssertEqualIntegers(t, response.Code, http.StatusNotFound)
	})
}

func TestPostPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{}, []string{}, nil,
	}
	cases := []struct {
		player string
		score  int
	}{
		{"Angelina", 20},
		{"Floyd", 10},
	}
	server, _ := NewPlayerServer(&store)
	for _, c := range cases {
		t.Run(testName(c.player), func(t *testing.T) {
			request := postRequest(c.player)
			response := httptest.NewRecorder()
			server.ServeHTTP(response, request)
			AssertEqualIntegers(t, response.Code, http.StatusAccepted)
			AssertPlayerInStore(t, store.winCalls, c)
		})
	}
}

func getRequest(player string) (*http.Request) {
	request, _ := http.NewRequest(http.MethodGet, getEndpoint(player), nil)
	return request
}

func postRequest(player string) (*http.Request) {
	request, _ := http.NewRequest(http.MethodPost, getEndpoint(player), nil)
	return request
}

func testName(player string) string {
	return fmt.Sprintf("returns %s's score", player)
}

func getEndpoint(player string) string {
	return "/players/" + player
}
