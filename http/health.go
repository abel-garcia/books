package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /health [get]
func (h Handler) HealthCheck() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		res := map[string]interface{}{
			"data": "Server is up and running",
		}

		ctx.JSON(http.StatusOK, res)
	}
}
