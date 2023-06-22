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

var artistCollection *mongo.Collection = configs.GetCollection(configs.DB, "artists")

func CreateArtist(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var artist models.Artist
	defer cancel()

	if err := c.BodyParser(&artist); err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			responses.Response{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    &fiber.Map{"data": err.Error()},
			})
	}

	newArtist := models.Artist{
		Id:               primitive.NewObjectID(),
		Name:             artist.Name,
		ArtistSince:      artist.ArtistSince,
		Genres:           artist.Genres,
		MonthlyListeners: artist.MonthlyListeners,
		Website:          artist.Website,
	}

	result, err := artistCollection.InsertOne(ctx, newArtist)
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

func GetArtist(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	artistId := c.Params("artistId")
	var artist models.Artist
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(artistId)

	err := artistCollection.FindOne(ctx, models.Artist{Id: objId}).Decode(&artist)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			responses.Response{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    &fiber.Map{"data": err.Error()},
			})
	}

	if !strings.Contains(artist.Id.String(), artistId) {
		return c.Status(http.StatusNotFound).JSON(
			responses.Response{
				Status:  http.StatusNotFound,
				Message: "error",
				Data:    &fiber.Map{"data": "Artist not found"},
			})
	}

	return c.Status(http.StatusOK).JSON(
		responses.Response{
			Status:  http.StatusOK,
			Message: "success",
			Data:    &fiber.Map{"data": artist},
		})
}

func EditArtist(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	artistId := c.Params("artistId")
	var artist models.Artist
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(artistId)

	if err := c.BodyParser(&artist); err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			responses.Response{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    &fiber.Map{"data": err.Error()},
			})
	}

	update := bson.M{
		"name":              artist.Name,
		"artist_since":      artist.ArtistSince,
		"genres":            artist.Genres,
		"monthly_listeners": artist.MonthlyListeners,
		"website":           artist.Website,
	}

	result, err := artistCollection.UpdateOne(ctx, models.Artist{Id: objId}, bson.M{"$set": update})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			responses.Response{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    &fiber.Map{"data": err.Error()},
			})
	}

	//	var updatedArtist models.Artist
	//	if result.MatchedCount == 1 {
	//		err := artistCollection.FindOne(ctx, models.Artist{Id: objId}).Decode(&updatedArtist)
	//
	//		if err != nil {
	//			return c.Status(http.StatusInternalServerError).JSON(
	//				responses.Response{
	//					Status:  http.StatusInternalServerError,
	//					Message: "error",
	//					Data:    &fiber.Map{"data": err.Error()},
	//				})
	//		}
	//	}

	return c.Status(http.StatusOK).JSON(
		responses.Response{
			Status:  http.StatusOK,
			Message: "success",
			Data:    &fiber.Map{"data": result},
		})
}

func DeleteArtist(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	artistId := c.Params("artistId")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(artistId)

	result, err := artistCollection.DeleteOne(ctx, models.Artist{Id: objId})
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
				Data:    &fiber.Map{"data": "Artist not found"},
			})
	}

	return c.Status(http.StatusOK).JSON(
		responses.Response{
			Status:  http.StatusOK,
			Message: "success",
			Data:    &fiber.Map{"data": "Artist deleted"},
		})
}

func GetArtists(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var artists []models.Artist
	defer cancel()

	results, err := artistCollection.Find(ctx, bson.M{})

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
		var artist models.Artist
		if err = results.Decode(&artist); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(
				responses.Response{
					Status:  http.StatusInternalServerError,
					Message: "error",
					Data:    &fiber.Map{"data": err.Error()},
				})
		}
		artists = append(artists, artist)
	}

	return c.Status(http.StatusOK).JSON(
		responses.Response{
			Status:  http.StatusOK,
			Message: "success",
			Data:    &fiber.Map{"data": artists},
		})
}
