package api

import (
	"github.com/gin-gonic/gin"
)

type (
	API struct {
		Router  *gin.Engine
		Address string
	}
)

// New returns a new instance of API.
//
// It creates a new gin router and adds the routes.
// Returns a pointer to the API struct.
func New() *API {
	r := gin.Default()

	r.GET("/ping", pingHandler)
	r.GET("/version", versionHandler)

	return &API{
		Router: r,
	}
}

// Run runs the API server.
//
// It returns an error if the underlying router fails to run.
func (a *API) Run() error {
	return a.Router.Run(a.Address)
}
