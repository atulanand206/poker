package poker

import (
	"github.com/gorilla/websocket"
	"net/http"
	"fmt"
	"encoding/json"
	"html/template"
)

const (
	contentTypeKey             = "content-type"
	contentTypeApplicationJson = "application/json"
	htmlTemplatePath           = "game.html"
)

type Player struct {
	Name string
	Wins int
}

type PlayerStore interface {
	getPlayerScore(player string) int
	RecordWin(player string)
	GetLeague() League
}

type PlayerServer struct {
	store    PlayerStore
	http.Handler
	template *template.Template
}

func NewPlayerServer(store PlayerStore) (*PlayerServer, error) {
	p := new(PlayerServer)
	tmpl, err := template.ParseFiles("game.html")
	if err != nil {
		return nil, fmt.Errorf("problem opening %s %v", htmlTemplatePath, err)
	}
	p.template = tmpl
	p.store = store
	p.Handler = initializeRoutes(p)
	return p, nil
}

func initializeRoutes(p *PlayerServer) *http.ServeMux {
	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playerHandler))
	router.Handle("/game", http.HandlerFunc(p.gameHandler))
	router.Handle("/ws", http.HandlerFunc(p.webSocket))
	return router
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (server *PlayerServer) webSocket(writer http.ResponseWriter, request *http.Request) {
	conn, _ := upgrader.Upgrade(writer, request, nil)
	_, winnerMsg, _ := conn.ReadMessage()
	server.store.RecordWin(string(winnerMsg))
}

func (server *PlayerServer) gameHandler(writer http.ResponseWriter, request *http.Request) {
	server.template.Execute(writer, nil)
}

func (server *PlayerServer) leagueHandler(writer http.ResponseWriter, request *http.Request) {
	json.NewEncoder(writer).Encode(server.getLeagueTable())
	writer.Header().Set(contentTypeKey, contentTypeApplicationJson)
}

func (server *PlayerServer) getLeagueTable() League {
	return server.store.GetLeague()
}

func (server *PlayerServer) playerHandler(writer http.ResponseWriter, request *http.Request) {
	player := playerNameFromRequest(*request)
	switch request.Method {
	case http.MethodPost:
		server.processWin(writer, player)
	case http.MethodGet:
		server.showScore(writer, player)
	}
}

func (server *PlayerServer) processWin(writer http.ResponseWriter, player string) {
	server.store.RecordWin(player)
	writer.WriteHeader(http.StatusAccepted)
}

func (server *PlayerServer) showScore(writer http.ResponseWriter, player string) {
	score := server.store.getPlayerScore(player)
	if score == 0 {
		writer.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(writer, score)
}

func playerNameFromRequest(request http.Request) string {
	return request.URL.Path[len("/players/"):]
}

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   League
}

func (s *StubPlayerStore) getPlayerScore(player string) int {
	score := s.scores[player]
	return score
}

func (s *StubPlayerStore) RecordWin(player string) {
	s.winCalls = append(s.winCalls, player)
}

func (s *StubPlayerStore) GetLeague() League {
	return s.league
}