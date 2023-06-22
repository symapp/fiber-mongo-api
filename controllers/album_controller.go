package controllers

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/symapp/fiber-mongo-api/configs"
	"github.com/symapp/fiber-mongo-api/models"
	"github.com/symapp/fiber-mongo-api/responses"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var albumCollection *mongo.Collection = configs.GetCollection(configs.DB, "albums")

func CreateAlbum(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var album models.Album
	defer cancel()

	if err := c.BodyParser(&album); err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			responses.Response{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    &fiber.Map{"data": err.Error()},
			})
	}

	newAlbum := models.Album{
		Id:          primitive.NewObjectID(),
		Name:        album.Name,
		Length:      album.Length,
		ArtistId:    album.ArtistId,
		AmtSongs:    album.AmtSongs,
		ReleaseYear: album.ReleaseYear,
		Sales:       album.Sales,
		Producer:    album.Producer,
		Genre:       album.Genre,
		Label:       album.Label,
		Studio:      album.Studio,
	}

	result, err := albumCollection.InsertOne(ctx, newAlbum)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			responses.Response{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    &fiber.Map{"data": err.Error()},
			})
	}

	return c.Status(http.StatusCreated).JSON(
		responses.Response{
			Status:  http.StatusCreated,
			Message: "success",
			Data:    &fiber.Map{"data": result},
		})
}

func GetAlbum(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	albumId := c.Params("albumId")
	var album models.Album
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(albumId)

	err := albumCollection.FindOne(ctx, models.Album{Id: objId}).Decode(&album)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			responses.Response{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    &fiber.Map{"data": err.Error()},
			})
	}

	if !strings.Contains(album.Id.String(), albumId) {
		return c.Status(http.StatusNotFound).JSON(
			responses.Response{
				Status:  http.StatusNotFound,
				Message: "error",
				Data:    &fiber.Map{"data": "Album not found"},
			})
	}

	return c.Status(http.StatusOK).JSON(
		responses.Response{
			Status:  http.StatusOK,
			Message: "success",
			Data:    &fiber.Map{"data": album},
		})
}

func EditAlbum(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	albumId := c.Params("albumId")
	var album models.Album
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(albumId)

	if err := c.BodyParser(&album); err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			responses.Response{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    &fiber.Map{"data": err.Error()},
			})
	}

	result, err := albumCollection.UpdateOne(ctx, models.Album{Id: objId}, bson.M{"$set": album})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			responses.Response{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    &fiber.Map{"data": err.Error()},
			})
	}

	return c.Status(http.StatusOK).JSON(
		responses.Response{
			Status:  http.StatusOK,
			Message: "success",
			Data:    &fiber.Map{"data": result},
		})
}

func DeleteAlbum(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	albumId := c.Params("albumId")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(albumId)
	result, err := albumCollection.DeleteOne(ctx, models.Album{Id: objId})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			responses.Response{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    &fiber.Map{"data": err.Error()},
			})
	}

	if result.DeletedCount < 1 {
		return c.Status(http.StatusNotFound).JSON(
			responses.Response{
				Status:  http.StatusNotFound,
				Message: "error",
				Data:    &fiber.Map{"data": "Album not found"},
			})
	}

	return c.Status(http.StatusOK).JSON(
		responses.Response{
			Status:  http.StatusOK,
			Message: "success",
			Data:    &fiber.Map{"data": "Album deleted"},
		})
}

func GetAlbums(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var albums []models.Album
	artistId := c.Query("artistId", "")
	defer cancel()

	filter := models.Album{}

	if artistId != "" {
		filter = models.Album{ArtistId: artistId}
	}

	results, err := albumCollection.Find(ctx, filter)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			responses.Response{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    &fiber.Map{"data": err.Error()},
			})
	}

	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleAlbum models.Album
		if err = results.Decode(&singleAlbum); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(
				responses.Response{
					Status:  http.StatusInternalServerError,
					Message: "error",
					Data:    &fiber.Map{"data": err.Error()},
				})
		}
		albums = append(albums, singleAlbum)
	}

	return c.Status(http.StatusOK).JSON(
		responses.Response{
			Status:  http.StatusOK,
			Message: "success",
			Data:    &fiber.Map{"data": albums},
		})
}
