package di

import (
	"github.com/FlowingSPDG/ProjectDeus/server/controller/cs2loghandler"
	"github.com/FlowingSPDG/ProjectDeus/server/controller/gin"
)

// TODO: 実際のgateway/controller/usecaseの実装を返すようにする
func InitializeLogController() cs2loghandler.LogController {
	return cs2loghandler.NewLogController(nil)
}

func InitializeRegisterServerController() gin.RegisterServerController {
	return gin.NewRegisterServerController(nil)
}
