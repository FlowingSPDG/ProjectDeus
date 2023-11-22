package database

import (
	"github.com/FlowingSPDG/ProjectDeus/server/entity"
)

// https://github.com/FlowingSPDG/get5loader/blob/main/backend/gateway/database/database.go

type GameServerRepository interface {
	// RegisterGameServer registers a game server in the database
	RegisterGameServer(userID entity.UserID, name string, public bool, ip string, rcon string) (entity.GameServerID, error)
	// GetGameServer returns a game server from the database
	GetGameServer(gameServerID entity.GameServerID) (*entity.GameServer, error)
	// VerifyGameServer verifies a game server in the database
	VerifyGameServer(gameServerID entity.GameServerID) error
}
