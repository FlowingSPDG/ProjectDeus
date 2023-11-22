package di

import (
	"github.com/FlowingSPDG/ProjectDeus/server/controller/cs2loghandler"
	"github.com/FlowingSPDG/ProjectDeus/server/controller/gin"
	"github.com/FlowingSPDG/ProjectDeus/server/gateway/database/local"
	"github.com/FlowingSPDG/ProjectDeus/server/service/servercfg"
	"github.com/FlowingSPDG/ProjectDeus/server/usecase"
)

func InitializeLogController() cs2loghandler.LogController {
	db := local.NewInmemoryGameServerRepository()
	cfg := servercfg.NewServerCfg("http://localhost:3090/servers/%s/log")
	uc := usecase.NewGameServerUsecase(db, cfg)
	return cs2loghandler.NewLogController(uc)
}

func InitializeRegisterServerController() gin.RegisterServerController {
	db := local.NewInmemoryGameServerRepository()
	cfg := servercfg.NewServerCfg("http://localhost:3090/servers/%s/log")
	uc := usecase.NewGameServerUsecase(db, cfg)
	return gin.NewRegisterServerController(uc)
}
