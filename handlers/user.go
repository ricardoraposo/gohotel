package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/ricardoraposo/gohotel/data"
	"github.com/ricardoraposo/gohotel/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := h.userStore.GetUserById(c.Context(), id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.JSON(map[string]string{"message": "user not found"})
		}
		return err
	}
	return c.JSON(user)
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.userStore.DeleteUser(c.Context(), id); err != nil {
		return err
	}
	return c.JSON(map[string]string{"deleted": id})
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var update data.UpdateUserRequest
	if err := c.BodyParser(&update); err != nil {
		return err
	}

	if err := h.userStore.UpdateUser(c.Context(), id, update); err != nil {
		return err
	}

	return c.JSON(map[string]string{"updated": id})
}

func (h *UserHandler) PostUser(c *fiber.Ctx) error {
	var params data.CreateUserRequest
	if err := c.BodyParser(&params); err != nil {
		return err
	}
	if err := params.Validate(); err != nil {
		return c.JSON(err)
	}
	user, err := data.NewUserFromRequest(params)
	if err != nil {
		return err
	}
	createdUser, err := h.userStore.InsertUser(c.Context(), user)
	if err != nil {
		return err
	}
	return c.JSON(createdUser)
}

func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	users, err := h.userStore.GetUsers(c.Context())
	if err != nil {
		return err
	}
	return c.JSON(users)
}
