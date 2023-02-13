package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PingController struct{}

func NewPingController() *PingController {
	return &PingController{}
}

// @Summary	Ping
// @Tags Ping
// @Description	Do a ping request to verifify if server is responding
// @Produce	json
// @Response	200	{object} TextResponse
// @Router	/ping [get]
func (pc *PingController) Pong(c *gin.Context) {
	c.JSON(http.StatusOK, NewTextResponse("pong"))
}
