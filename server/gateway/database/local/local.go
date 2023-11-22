package local

import (
	"github.com/FlowingSPDG/ProjectDeus/server/entity"
	"github.com/FlowingSPDG/ProjectDeus/server/gateway/database"
	"github.com/FlowingSPDG/ProjectDeus/server/service/uuid"
)

// in-memory database system for testing purpose.
// DO NOT USE THIS IN PRODUCTION.

type inmemoryGameServerRepository struct {
	uuidGenerator uuid.UUIDGenerator
	gameServers   map[entity.GameServerID]*entity.GameServer
}

// NewInmemoryGameServerRepository returns a new in-memory database system for testing purpose.
// This should not be used in production.
func NewInmemoryGameServerRepository() database.GameServerRepository {
	return &inmemoryGameServerRepository{
		uuidGenerator: uuid.NewUUIDGenerator(), // テスト用途なのでDIはしない
		gameServers:   map[entity.GameServerID]*entity.GameServer{},
	}
}

// RegisterGameServer implements database.GameServerRepository.
func (im *inmemoryGameServerRepository) RegisterGameServer(userID entity.UserID, name string, public bool, ip string, rcon string) (entity.GameServerID, error) {
	id := im.uuidGenerator.Generate()
	im.gameServers[entity.GameServerID(id)] = &entity.GameServer{
		ID:           entity.GameServerID(id),
		Owner:        userID,
		Name:         name,
		Public:       public,
		IP:           ip,
		InUse:        false,
		IsVeritifed:  false,
		RconPassword: rcon,
	}
	return entity.GameServerID(id), nil
}

// GetGameServer implements database.GameServerRepository.
func (im *inmemoryGameServerRepository) GetGameServer(gameServerID entity.GameServerID) (*entity.GameServer, error) {
	gs, ok := im.gameServers[gameServerID]
	if !ok {
		return nil, database.ErrNotFound
	}
	return gs, nil
}

// VerifyGameServer implements database.GameServerRepository.
func (im *inmemoryGameServerRepository) VerifyGameServer(gameServerID entity.GameServerID) error {
	gs, ok := im.gameServers[gameServerID]
	if !ok {
		return database.ErrNotFound
	}
	gs.IsVeritifed = true
	return nil
}
