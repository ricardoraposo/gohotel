package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ricardoraposo/gohotel/data"
)

func AdminAuth(c *fiber.Ctx) error {
	user, ok := c.Context().UserValue("user").(*data.User)
    fmt.Println(user)
	if !ok {
		return fmt.Errorf("Not authorized")
	}
	if !user.IsAdmin {
		return fmt.Errorf("Not enough credentials")
	}
    return c.Next()
}
