package main

import (
	"github.com/IBM/go-gin-app/routers"
	// "goginapp/plugins" if you create your own plugins import them here
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"os"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}

func main() {

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	router := gin.Default()
	router.RedirectTrailingSlash = false

	// Swagger endpoints
	url := ginSwagger.URL("/swagger.yaml") // The url pointing to API definition
	router.GET("/swagger/api-docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	router.GET("/swagger/api-docs", routers.SwaggerDocRedirect)

	// Static webpage content endpoints
	router.GET("/swagger.yaml", routers.SwaggerAPI)
	router.LoadHTMLGlob("public/*.html")
	router.Use(static.Serve("/", static.LocalFile("./public", false)))
	router.GET("/", routers.Index)
	router.NoRoute(routers.NotFoundError)
	router.GET("/500", routers.InternalServerError)

	// Health endpoint
	router.GET("/health", routers.HealthGET)

	log.Info("Starting goginapp on port " + port())

	router.Run(port())
}
