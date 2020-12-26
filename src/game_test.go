package poker

import (
	"github.com/gorilla/websocket"
	"testing"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"
)

func TestGame(t *testing.T) {
	t.Run("GET /game returns 200", func(t *testing.T) {
		request := getGameRequest()
		response := httptest.NewRecorder()
		server, _ := NewPlayerServer(&StubPlayerStore{})
		server.ServeHTTP(response, request)
		AssertStatus(t, response, http.StatusOK)
	})

	t.Run("when we get a message from websocket it is a winner", func(t *testing.T) {
		store := &StubPlayerStore{}
		serve := mustMakePlayerServer(store, t)
		server := httptest.NewServer(serve)
		defer server.Close()

		wsUrl := "ws" + strings.TrimPrefix(server.URL, "http") + "/ws"
		ws := mustMakeWebSocket(t, wsUrl)
		defer ws.Close()

		winner := "Macy"
		writeWebSocketMessage(t, ws, winner)

		time.Sleep(10 * time.Millisecond)
		AssertPlayerWin(t, store, winner)
	})
}

func writeWebSocketMessage(t *testing.T, ws *websocket.Conn, winner string) {
	t.Helper()
	if err := ws.WriteMessage(websocket.TextMessage, []byte(winner)); err != nil {
		t.Fatalf("could not send message over ws connection %v", err)
	}
}

func mustMakeWebSocket(t *testing.T, wsUrl string) *websocket.Conn {
	ws, _, err := websocket.DefaultDialer.Dial(wsUrl, nil)
	if err != nil {
		t.Fatalf("could not open a ws connection on %s %v", wsUrl, err)
	}
	return ws
}

func mustMakePlayerServer(store *StubPlayerStore, t *testing.T) (*PlayerServer) {
	serve, err := NewPlayerServer(store)
	if err != nil {
		t.Fatal("problem creating player server", err)
	}
	return serve
}

func getGameRequest() *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/game", nil)
	return request
}
