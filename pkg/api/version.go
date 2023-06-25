package api

import (
	"net/http"

	"github.com/danstis/CrowdSound/internal/version"
	"github.com/gin-gonic/gin"
)

func versionHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"version": version.Version,
	})
}
