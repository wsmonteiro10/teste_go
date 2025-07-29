package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping retorna pong
// @Summary Retorna pong
// @Description Endpoint básico para teste de vida do servidor.
// @Tags health
// @Produce json
// @Success 200 {object} map[string]string
// @Router /ping [get]
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
