package playplay

import (
	"net/http"
	"time"
)

const (
	baseURL            = "https://www.playplay.io/api"
	statusPath         = "/status"
	usersPath          = "/users"
	challengesPath     = "/challenges"
	matchesPath        = "/matches"
	sessionsPath       = "/sessions"
	currentSessionPath = sessionsPath + "/current"
	teamsPath          = "/teams"
	gamesPath          = "/games"
	subscriptionsPath  = "/subscriptions"
	challengePath      = "/challenge"
)

/*
{"_links":{
	"self":{"href":"https://www.playplay.io/api/"},
	COMPLETE - "status":{"href":"https://www.playplay.io/api/status"},
	"users":{"href":"https://www.playplay.io/api/users/{?cursor,size,sort,offset,total_count,team_id,captain}","templated":true},
	"challenges":{"href":"https://www.playplay.io/api/challenges/{?cursor,size,sort,offset,total_count,team_id}","templated":true},
	"matches":{"href":"https://www.playplay.io/api/matches/{?cursor,size,sort,offset,total_count,team_id}","templated":true},
	"current_season":{"href":"https://www.playplay.io/api/seasons/current/{?team_id}","templated":true},
	"seasons":{"href":"https://www.playplay.io/api/seasons/{?cursor,size,sort,offset,total_count,team_id}","templated":true},
	"teams":{"href":"https://www.playplay.io/api/teams/{?cursor,size,sort,offset,total_count,active,game_id}","templated":true},
	"games":{"href":"https://www.playplay.io/api/games/{?cursor,size,sort,offset,total_count}","templated":true},
	"subscriptions":{"href":"https://www.playplay.io/api/subscriptions"},
	"challenge":{"href":"https://www.playplay.io/api/challenges/{id}","templated":true},
	"match":{"href":"https://www.playplay.io/api/matches/{id}","templated":true},
	"user":{"href":"https://www.playplay.io/api/users/{id}","templated":true},
	"season":{"href":"https://www.playplay.io/api/seasons/{id}","templated":true},
	"team":{"href":"https://www.playplay.io/api/teams/{id}","templated":true},
	"game":{"href":"https://www.playplay.io/api/games/{id}","templated":true}}
}
*/

// PlayPlay
type PlayPlay struct {
	hc    http.Client
	Debug bool
}

// NewPlayPlay
func NewPlayPlay() *PlayPlay {
	p := &PlayPlay{
		hc: http.Client{
			Timeout: time.Duration(10 * time.Second),
		},
		Debug: false,
	}
	return p
}
