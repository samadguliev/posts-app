package handlers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/samadguliev/posts-app/backend/db"
	"github.com/samadguliev/posts-app/backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func PostList(c *fiber.Ctx) error {
	var posts []models.Post

	var collection = db.Client.Database("posts").Collection("posts")
	cursor, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err = cursor.All(context.TODO(), &posts); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	postAPIList := make([]models.PostAPI, len(posts))

	for key, postAPI := range postAPIList {
		postAPIList[key] = postAPI.GeneratePostApi(posts[key])
	}

	return c.Status(200).JSON(postAPIList)
}

func AddPost(c *fiber.Ctx) error {
	post := new(models.Post)
	post.CreatedAt = time.Now()

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
