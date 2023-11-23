package usecase

import (
	"github.com/FlowingSPDG/ProjectDeus/server/entity"
	"github.com/FlowingSPDG/ProjectDeus/server/gateway/database"
	"github.com/FlowingSPDG/ProjectDeus/server/service/servercfg"
)

type GameServerUsecase interface {
	// RegisterGameServer registers a temporary game server in the database.
	// 返り値: server.cfg に追記する文字列
	// TODO: JSON化する
	RegisterGameServer(userDiscordID entity.UserDiscordID, name string, public bool, ip string, rcon string) (string, error)
	// VerifyGameServer verifies a game server in the database. This should be called when the API server receives a log from the game server.
	VerifyGameServer(gameServerID entity.GameServerID, ip string) error
}

type gameServerUsecase struct {
	gameServerRepository database.GameServerRepository
	serverCfgGenerator   servercfg.ServerCfg
}

func NewGameServerUsecase(
	gameServerRepository database.GameServerRepository,
	serverCfgGenerator servercfg.ServerCfg,
) GameServerUsecase {
	return &gameServerUsecase{
		gameServerRepository: gameServerRepository,
		serverCfgGenerator:   serverCfgGenerator,
	}
}

// RegisterGameServer implements GameServerUsecase.
func (uc *gameServerUsecase) RegisterGameServer(userDiscordID entity.UserDiscordID, name string, public bool, ip string, rcon string) (string, error) {
	serverID, err := uc.gameServerRepository.RegisterGameServer(userDiscordID, name, public, ip, rcon)
	if err != nil {
		return "", err
	}

	return uc.serverCfgGenerator.GenerateLogAddressHTTP(serverID), nil
}

// VerifyGameServer implements GameServerUsecase.
func (uc *gameServerUsecase) VerifyGameServer(gameServerID entity.GameServerID, ip string) error {
	gs, err := uc.gameServerRepository.GetGameServer(gameServerID)
	if err != nil {
		return err
	}

	// ゲームサーバーのIPではない場合、不正なリクエストとして処理する
	if gs.IP != ip {
		return database.ErrProhibited
	}

	// 問題ないので、サーバーを認証済みとしてマーク
	return uc.gameServerRepository.VerifyGameServer(gs.ID)
}
