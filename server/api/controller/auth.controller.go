package controller

import (
	"scsystem/internal/schema"
	"scsystem/internal/service"
	"scsystem/pkg/validator"

	"github.com/gofiber/fiber/v2"
)

// SignUp
// @Description Register account for admin.
// @Tags auth
// @Accept json
// @Produce json
// @Param sign_up body schema.SignUpRequest true "Sign Up"
// @Success 200 {object} schema.SignUpResponse
// @Router /auth/signup [POST]
func SignUp(c *fiber.Ctx) error {
	var authService = service.NewAuth()
	var req schema.SignUpRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(schema.Error{
			Success: false,
			Msg:     err.Error(),
		})
	}
	if _validate := validator.Validator(req); _validate != "" {
		return c.Status(fiber.StatusBadRequest).JSON(schema.Error{
			Success: false,
			Msg:     _validate,
		})
	}
	if err := authService.SignUp(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(schema.Error{
			Success: false,
			Msg:     err.Error(),
		})
	}
	return c.JSON(schema.SignUpResponse{
		Success: true,
		Msg:     "your account was created successfully",
	})
}

// SignIn
// @Description Sign in account for admin.
// @Tags auth
// @Accept json
// @Produce json
// @Param sign_in body schema.SignInRequest true "Sign In"
// @Success 200 {object} schema.SignInResponse
// @Router /auth/signin [POST]
func SignIn(c *fiber.Ctx) error {
	var authService = service.NewAuth()
	var req schema.SignInRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(schema.Error{
			Success: false,
			Msg:     err.Error(),
		})
	}
	if _validate := validator.Validator(req); _validate != "" {
		return c.Status(fiber.StatusBadRequest).JSON(schema.Error{
			Success: false,
			Msg:     _validate,
		})
	}
	token, err := authService.SignIn(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(schema.Error{
			Success: false,
			Msg:     "email or password incorrect, please try again",
		})
	}
	return c.JSON(schema.SignInResponse{
		Success: true,
		Token:   token,
		Email:   req.Email,
	})
}

// GetMe
// @Description Get information by student id
// @Tags user
// @Accept json
// @Produce json
// @Param sid query string true "student id"
// @Success 200 {object} schema.SignInResponse
// @Router /user [GET]
func GetMe(c *fiber.Ctx) error {
	var authService = service.NewAuth()
	sid := c.Query("sid", "")
	if sid == "" {
		return c.Status(fiber.StatusBadRequest).JSON(schema.Error{
			Success: false,
			Msg:     "missing query params sid",
		})
	}
	data, err := authService.GetMe(sid)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(schema.Error{
			Success: false,
			Msg:     "student id invalid",
		})
	}
	return c.JSON(data)
}
