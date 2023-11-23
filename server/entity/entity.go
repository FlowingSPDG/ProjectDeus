package entity

import "time"

type UserID string
type UserDiscordID string
type MatchID string
type PlayerID string
type GameServerID string

type GameServer struct {
	ID           GameServerID
	Owner        UserDiscordID
	Name         string
	Public       bool
	IP           string
	InUse        bool
	IsVeritifed  bool
	RconPassword string
}

type MatchState int

const (
	MatchStateUnknown   MatchState = iota // Unknown
	MatchStateStandby                     // Match created and loaded to game server
	MatchStateWarmup                      // Match is in warmup, players can join
	MatchStateLive                        // Match is live
	MatchStateEnded                       // Match has ended
	MatchStateCancelled                   // Match has been cancelled
	MatchStateError                       // Match has errored, should be investigated by admin
)

type Match struct {
	ID         MatchID
	Owner      UserDiscordID
	Server     GameServerID
	StartedAt  *time.Time
	EndedAt    *time.Time
	IsPug      bool
	Players    []PlayerID
	State      MatchState
	Team1Score int
	Team2Score int
	Map        string
	Winner     string
}

type User struct {
	// ID        UserID // should not be used since ProjectDeus is using Discord as the main auth provider for now
	DiscordID UserDiscordID
	// SteamID  string // SteamID64, should be used for game stats and in-game admin calls?
	IsAdmin bool
}
