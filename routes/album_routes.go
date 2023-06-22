package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/symapp/fiber-mongo-api/controllers"
)

func AlbumRoute(app *fiber.App) {
	app.Post("/api/albums", controllers.CreateAlbum)
	app.Get("/api/albums/:albumId", controllers.GetAlbum)
	app.Put("/api/albums/:albumId", controllers.EditAlbum)
	app.Delete("/api/albums/:albumId", controllers.DeleteAlbum)
	app.Get("/api/albums", controllers.GetAlbums)

}
