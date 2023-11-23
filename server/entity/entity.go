package entity

import "time"

type UserID string
type PlayerID string
type GameServerID string

type GameServer struct {
	ID           GameServerID
	Owner        UserID
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
	ID        string
	Owner     UserID
	Server    GameServerID
	StartedAt *time.Time
	EndedAt   *time.Time
	IsPug     bool
	Players   []PlayerID
	State     MatchState
}
