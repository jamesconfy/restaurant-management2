package main

import "restaurant-management/cmd"

func main() {
	// @title           Project
	// @version         1.0
	// @description     Server for {{your server link}}
	// @termsOfService  http://swagger.io/terms/

	// @contact.name   Confidence James
	// @contact.url    http://github.com/jamesconfy
	// @contact.email  bobdence@gmail.com

	// @license.name  Apache 2.0
	// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

	// @host      localhost:80
	// @schemes http https
	// @BasePath  /v1

	// @securityDefinitions.apiKey  ApiKeyAuth
	// @in header
	// @name Authorisation
	cmd.Setup()
}
