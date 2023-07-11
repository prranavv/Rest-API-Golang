package handlers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/prranavv/Rest-API-Golang/database"
)

type libraryDTO struct {
	Name    string `json:"name" bson:"name"`
	Address string `json:"address" bson:"address"`
}

func CreateLibrary(c *fiber.Ctx) error {
	nlibrary := new(libraryDTO)
	if err := c.BodyParser(nlibrary); err != nil {
		return err
	}
	librarycollection := database.GetCollection("libraries")
	ndoc, err := librarycollection.InsertOne(context.TODO(), nlibrary)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"id": ndoc.InsertedID})
}
