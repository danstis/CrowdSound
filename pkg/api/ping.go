package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Ping returns a simple 204 response if the api is online.
// @Schemes http https
// @Tags utility
// @Success 204
// @Router /ping [get]
func pingHandler(c *gin.Context) {
	c.JSON(http.StatusNoContent, nil)
}
