package main

import (
	_ "coffee-layered-architecture/docs"
	"coffee-layered-architecture/internal/app"
)

const configPath = "configs/main"

// @title VTB Hack API
// @version 1.0
// @description API for VTB Hack
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api

// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	app.Run(configPath)
}
