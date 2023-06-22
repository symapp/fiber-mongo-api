package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/symapp/fiber-mongo-api/configs"
	"github.com/symapp/fiber-mongo-api/routes"
)

func main() {
	app := fiber.New()

	configs.ConnectDB()

	//routes.UserRoute(app)
	routes.ArtistRoute(app)
	routes.AlbumRoute(app)

	// CORS middleware configuration
	corsConfig := cors.Config{
		AllowOrigins:                      "http://localhost:4200",
		AllowMethods:                      "GET,POST,PUT,DELETE",
		AllowHeaders:                      "Origin, Content-Type, Accept, Accept-Language, Content-Length",
		AllowCredentials:                  true,
	}

	// CORS middleware
	app.Use(cors.New(corsConfig))

	app.Listen(":4201")
}
