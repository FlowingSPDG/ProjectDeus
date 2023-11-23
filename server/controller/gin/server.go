package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/FlowingSPDG/ProjectDeus/server/usecase"
)

type RegisterServerController interface {
	Handle(c *gin.Context)
}

type registerServerController struct {
	uc usecase.GameServerUsecase
}

func NewRegisterServerController(
	uc usecase.GameServerUsecase,
) RegisterServerController {
	return &registerServerController{
		uc: uc,
	}
}

type registerServerRequest struct {
	Name   string `json:"name"`
	Public bool   `json:"public"`
	IP     string `json:"ip"`
	RCON   string `json:"rcon"`
}

// Handle implements RegisterServerController.
func (rsc *registerServerController) Handle(c *gin.Context) {
	var req registerServerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	// TODO: ユーザーIDの取得をどうにかする
	cfg, err := rsc.uc.RegisterGameServer("FlowingSPDG_DEV", req.Name, req.Public, req.IP, req.RCON)
	if err != nil {
		panic(err)
	}

	c.String(http.StatusOK, cfg)
}
