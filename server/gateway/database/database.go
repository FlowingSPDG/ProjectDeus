package database

import (
	"github.com/FlowingSPDG/ProjectDeus/server/entity"
)

// https://github.com/FlowingSPDG/get5loader/blob/main/backend/gateway/database/database.go

type GameServerRepository interface {
	// RegisterGameServer registers a game server in the database
	RegisterGameServer(userDiscordID entity.UserDiscordID, name string, public bool, ip string, rcon string) (entity.GameServerID, error)
	// GetGameServer returns a game server from the database
	GetGameServer(gameServerID entity.GameServerID) (*entity.GameServer, error)
	// GetGameServers returns all game servers from the database
	GetGameServers() ([]*entity.GameServer, error)
	// GetAvailableGameServers returns all available/verified game servers from the database
	GetAvailableGameServers() ([]*entity.GameServer, error)
	// VerifyGameServer verifies a game server in the database
	VerifyGameServer(gameServerID entity.GameServerID) error
	// DeleteGameServer deletes a game server from the database
	DeleteGameServer(gameServerID entity.GameServerID) error
	// SetGameServerInUse sets the in use status of a game server in the database
	SetGameServerInUse(gameServerID entity.GameServerID, inUse bool) error
}

type CreateMatchInput struct {
	OwnerID  entity.UserDiscordID
	ServerID entity.GameServerID
	IsPug    bool
	Map      string
	Players  []entity.PlayerID
}

type MatchRepository interface {
	// CreateMatch creates a match in the database
	CreateMatch(input CreateMatchInput) (entity.MatchID, error)
	// GetMatch returns a match from the database
	GetMatch(matchID entity.MatchID) (*entity.Match, error)
	// StartMatch starts a match in the database
	StartMatch(matchID entity.MatchID) error
	// FinishMatch ends a match in the database
	FinishMatch(matchID entity.MatchID, winner string) error
	// UpdateScore updates the score of a match in the database
}
