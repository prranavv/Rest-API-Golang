package main

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/prranavv/Rest-API-Golang/database"
	"go.mongodb.org/mongo-driver/bson"
)

func initapp() error {
	err := loadEnv()
	if err != nil {
		return err
	}
	err = database.StartMongodb()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := initapp()
	if err != nil {
		panic(err)
	}
	defer database.CloseMongodb()
	app := fiber.New()

	app.Post("/", func(c *fiber.Ctx) error {
		sampledoc := bson.M{"name": "sample todo"}
		collection := database.GetCollection("todos")
		ndoc, err := collection.InsertOne(context.TODO(), sampledoc)
		if err != nil {
			return c.Status(400).SendString("Error inserting todo")
		}
		//send down info about todo
		return c.JSON(ndoc)

	})

	app.Listen(":3000")
}

func loadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	return nil
}
