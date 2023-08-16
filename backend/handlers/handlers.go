package handlers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/samadguliev/posts-app/backend/db"
	"github.com/samadguliev/posts-app/backend/models"
)

func PostList(c *fiber.Ctx) error {
	return c.SendStatus(200)
}

func AddPost(c *fiber.Ctx) error {
	post := new(models.Post)
	var collection = db.Client.Database("posts").Collection("posts")

	if err := c.BodyParser(post); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	_, err := collection.InsertOne(context.TODO(), post)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(post)
}
