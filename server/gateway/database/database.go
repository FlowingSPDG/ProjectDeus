package database

import (
	"github.com/FlowingSPDG/ProjectDeus/server/entity"
)

type GameServerRepository interface {
	// RegisterGameServer registers a game server in the database
	RegisterGameServer(userID entity.UserID, name string, public bool, ip string) (entity.GameServerID, error)
	// VerifyGameServer verifies a game server in the database
	VerifyGameServer(gameServerID entity.GameServerID) error
}
