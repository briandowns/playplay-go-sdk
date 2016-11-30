package playplay

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

// Auth
type Auth struct {
	Ok     bool   `json:"ok"`
	URL    string `json:"url"`
	Team   string `json:"team"`
	User   string `json:"user"`
	TeamID string `json:"team_id"`
	UserID string `json:"user_id"`
}

// Presence
type Presence struct {
	Ok              bool   `json:"ok"`
	Presence        string `json:"presence"`
	Online          bool   `json:"online"`
	AutoAway        bool   `json:"auto_away"`
	ManualAway      bool   `json:"manual_away"`
	ConnectionCount int    `json:"connection_count"`
	LastActivity    int    `json:"last_activity"`
}

// Ping
type Ping struct {
	Auth     Auth     `json:"auth"`
	Presence Presence `json:"presence"`
}

// Links
type Links struct {
	Self struct {
		Href string `json:"href"`
	} `json:"self"`
}

// Status
type Status struct {
	GamesCount int   `json:"games_count"`
	Games      Games `json:"games"`
	Pong       Pong  `json:"pong"`
	Links      Links `json:"_links"`
}

// Games
type Games struct {
	Chess     Chess     `json:"chess"`
	Ping      Ping      `json:"ping"`
	Pool      Pool      `json:"pool"`
	TicTacToe TicTacToe `json:"tictactoe"`
}

// BaseGame
type BaseGame struct {
	TeamsCount       int      `json:"teams_count"`
	ActiveTeamsCount int      `json:"active_teams_count"`
	APITeamsCount    int      `json:"api_teams_count"`
	UsersCount       int      `json:"users_count"`
	ChallengesCount  int      `json:"challenges_count"`
	MatchesCount     int      `json:"matches_count"`
	SeasonsCount     int      `json:"seasons_count"`
	Ping             Ping     `json:"ping"`
	Presence         Presence `json:"presence"`
}

// Chess
type Chess struct {
	BaseGame
}

// Pong
type Pong struct {
	BaseGame
}

// Pool
type Pool struct {
	BaseGame
}

// TicTacToe
type TicTacToe struct {
	BaseGame
}

// Status gets the current state of PlayPlay
func (p *PlayPlay) Status() (*Status, error) {
	req, err := http.NewRequest(http.MethodGet, baseURL+statusPath, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")

	res, err := p.hc.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if p.Debug {
		io.Copy(os.Stdout, res.Body)
	}

	var s Status
	if json.NewDecoder(res.Body).Decode(&s); err != nil {
		return nil, err
	}

	return &s, nil
}
