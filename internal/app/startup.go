package app

import (
	"golang-webapi-template/app/controllers"
	"golang-webapi-template/app/middlewares"
	"golang-webapi-template/core"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

var configuration *core.Configuration

func NewStartup(config core.Configuration) *core.Startup {
	instance := &core.Startup{}

	instance.ConfigureServicesFunc = ConfigurationServices
	instance.ConfigureFunc = Configuration

	configuration = &config

	return instance
}

func ConfigurationServices(container *core.Container) {
	// Add controller's constructor into container
	container.AddControllers(controllers.NewWeatherController)
}

func Configuration(engine *core.GinEngine) {
	// Use json logging
	engine.Use(middlewares.JsonLoggerMiddleware())
	engine.Use(gin.Recovery())

	// Add static file to routes
	engine.Use(static.Serve("/", static.LocalFile("./assets", false)))

	engine.UseControllers()
}