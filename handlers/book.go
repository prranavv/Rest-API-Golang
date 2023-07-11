package handlers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/prranavv/Rest-API-Golang/database"
	"github.com/prranavv/Rest-API-Golang/models"
	"go.mongodb.org/mongo-driver/bson"
)

type newbookDTO struct {
	Title     string `json:"title" bson:"title"`
	Author    string `json:"author" bson:"author"`
	ISBN      string `json:"isbn" bson:"isbn"`
	LibraryID string `json:"libraryid" bson:"libraryid"`
}

func Createbook(c *fiber.Ctx) error {
	createdata := new(newbookDTO)
	if err := c.BodyParser(createdata); err != nil {
		return err
	}
	//get the collection reference
	coll := database.GetCollection("libraries")
	//get the filter
	filter := bson.M{"_id": createdata.LibraryID}
	nbookdata := models.Book{
		Title:  createdata.Title,
		Author: createdata.Author,
		ISBN:   createdata.ISBN,
	}
	updatepayload := bson.M{"$push": bson.M{"books": nbookdata}}
	//update the library
	_, err := coll.UpdateOne(context.TODO(), filter, updatepayload)
	if err != nil {
		return err
	}

	return c.SendString("Book created successfully")
}
