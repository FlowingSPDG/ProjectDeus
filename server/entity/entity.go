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

type Match struct {
	ID        string
	Owner     UserID
	Server    GameServerID
	StartedAt *time.Time
	EndedAt   *time.Time
	IsPug     bool
	Players   []PlayerID
}
