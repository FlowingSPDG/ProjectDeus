package cs2loghandler

import (
	cs2log "github.com/janstuemmel/cs2-log"

	"github.com/FlowingSPDG/ProjectDeus/server/entity"
	"github.com/FlowingSPDG/ProjectDeus/server/usecase"
)

type LogController interface {
	Handle(ip string, id string, msg cs2log.Message) error
}

type logController struct {
	gameServerUsecase usecase.GameServerUsecase
}

func NewLogController(
	gameServerUsecase usecase.GameServerUsecase,
) LogController {
	return &logController{
		gameServerUsecase: gameServerUsecase,
	}
}

// Handle implements logController.
func (lh *logController) Handle(ip string, id string, msg cs2log.Message) error {
	return lh.gameServerUsecase.VerifyGameServer(entity.GameServerID(id), ip)
}
