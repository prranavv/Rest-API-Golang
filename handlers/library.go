package handlers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/prranavv/Rest-API-Golang/database"
	"github.com/prranavv/Rest-API-Golang/models"
	"go.mongodb.org/mongo-driver/bson"
)

// GET
func Getlibraries(c *fiber.Ctx) error {
	librarycoll := database.GetCollection("libraries")
	cursor, err := librarycoll.Find(context.TODO(), bson.M{})
	if err != nil {
		return err
	}
	var libraries []models.Library
	if err = cursor.All(context.TODO(), &libraries); err != nil {
		return err
	}
	return c.JSON(libraries)
}

type libraryDTO struct {
	Name    string   `json:"name" bson:"name"`
	Address string   `json:"address" bson:"address"`
	Empty   []string `json:"no_exists" bson:"books"`
}

// POST
func CreateLibrary(c *fiber.Ctx) error {
	nlibrary := new(libraryDTO)
	if err := c.BodyParser(nlibrary); err != nil {
		return err
	}
	nlibrary.Empty = make([]string, 0)
	librarycollection := database.GetCollection("libraries")
	ndoc, err := librarycollection.InsertOne(context.TODO(), nlibrary)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"id": ndoc.InsertedID})
}

// DELETE
func Deletelibrary(c *fiber.Ctx) error {
	id := c.Params("id")
	librarycoll := database.GetCollection("libraries")
	_, err := librarycoll.DeleteOne(context.TODO(), bson.M{"_`id": id})
	if err != nil {
		return err
	}
	return c.SendString("Library deleted successfully")
}
