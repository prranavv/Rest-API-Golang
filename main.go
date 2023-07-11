package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/prranavv/Rest-API-Golang/database"
	"github.com/prranavv/Rest-API-Golang/handlers"
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
	app := generateapp()
	//get the port from the env
	port := os.Getenv("PORT")
	app.Listen(":" + port)
}

func loadEnv() error {
	goenv := os.Getenv("GO_ENV")
	if goenv == "" {
		err := godotenv.Load()
		if err != nil {
			return err
		}
	}
	return nil
}

func generateapp() *fiber.App {
	app := fiber.New()
	//create health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	//create the library group and routes
	libgroup := app.Group("/library")
	libgroup.Get("/", handlers.Getlibraries)
	libgroup.Post("/", handlers.CreateLibrary)
	libgroup.Post("/book", handlers.Createbook)
	libgroup.Delete("/:id", handlers.Deletelibrary)
	return app
}
