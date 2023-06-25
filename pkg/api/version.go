package api

import (
	"net/http"

	"github.com/danstis/CrowdSound/internal/version"
	"github.com/gin-gonic/gin"
)

// @Summary Version returns the current version of the API.
// @Schemes http https
// @Description Returns the current version of the API.
// @Tags utility
// @Produce json
// @Success 200 {json} version
// @Router /version [get]
func versionHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"version": version.Version,
	})
}
