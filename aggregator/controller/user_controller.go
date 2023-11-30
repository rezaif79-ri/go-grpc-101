package controller

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/rezaif79-ri/go-grpc-101/aggregator/domain"
	"github.com/rezaif79-ri/go-grpc-101/aggregator/domain/construct"
)

type userController struct {
	domain.UserService
}

func NewUserController(userService *domain.UserService) domain.UserController {
	return &userController{
		UserService: *userService,
	}
}

func (uc *userController) GreetUser(c *fiber.Ctx) {
	type requestBody struct {
		FirstName  string `json:"first_name"`
		LastName   string `json:"last_name"`
		Salutation string `json:"salutation"`
	}
	var body requestBody
	err := c.BodyParser(&body)
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	res, err := uc.UserService.GreetUser(construct.GreetUserReq{
		Name:       strings.Join([]string{body.FirstName, body.LastName}, " "),
		Salutation: body.Salutation,
	})
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	c.Status(http.StatusOK).JSON(map[string]interface{}{
		"status":  http.StatusOK,
		"message": "OK",
		"data":    res,
	})

}
