package main

import "github.com/gofiber/fiber/v2"

const listenAddr = ":3000"
const dbUri = "mongodb://localhost:27017"
const userColl = "users"

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"message": err.Error()})
	},
}
