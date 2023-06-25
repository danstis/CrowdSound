package api

//go:generate swag init --parseDependency --parseInternal --output ../docs --dir ../../cmd/api

// API comment formats: https://github.com/swaggo/swag/blob/master/README.md#declarative-comments-format

import (
	"github.com/gin-gonic/gin"

	docs "github.com/danstis/CrowdSound/pkg/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Title = "CrowdSound API"

	r.GET("/ping", pingHandler)
	r.GET("/version", versionHandler)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

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
