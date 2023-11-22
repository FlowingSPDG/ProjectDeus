package di

import (
	"github.com/FlowingSPDG/ProjectDeus/server/controller/cs2loghandler"
	"github.com/FlowingSPDG/ProjectDeus/server/controller/gin"
	"github.com/FlowingSPDG/ProjectDeus/server/gateway/database/local"
	"github.com/FlowingSPDG/ProjectDeus/server/service/servercfg"
	"github.com/FlowingSPDG/ProjectDeus/server/usecase"
)

// localのインメモリDBが別インスタンスで生成されるので、データを参照できない場合がある
// テスト用途なので、いったん変数にキャプチャする

// TODO: 消す
// はやめに実際の外部に依存したDBを使うようにしたい
var uc usecase.GameServerUsecase

func init() {
	db := local.NewInmemoryGameServerRepository()
	cfg := servercfg.NewServerCfg("http://localhost:3090/servers/%s/log")
	uc = usecase.NewGameServerUsecase(db, cfg)
}

func InitializeLogController() cs2loghandler.LogController {
	return cs2loghandler.NewLogController(uc)
}

func InitializeRegisterServerController() gin.RegisterServerController {
	return gin.NewRegisterServerController(uc)
}
