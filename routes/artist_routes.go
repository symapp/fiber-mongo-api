package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/symapp/fiber-mongo-api/controllers"
)

func ArtistRoute(app *fiber.App) {
	app.Post("/api/artists", controllers.CreateArtist)
	app.Get("/api/artists/:artistId", controllers.GetArtist)
	app.Put("/api/artists/:artistId", controllers.EditArtist)
	app.Delete("/api/artists/:artistId", controllers.DeleteArtist)
	app.Get("/api/artists", controllers.GetArtists)
}
