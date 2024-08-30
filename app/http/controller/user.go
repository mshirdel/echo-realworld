package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mshirdel/echo-realworld/app/http/controller/dto"
	"github.com/mshirdel/echo-realworld/app/service"
	"github.com/mshirdel/echo-realworld/models"
)

type LoginUserRequest struct {
	User UserLogin `json:"user"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterUserRequest struct {
	User User `json:"user"`
}

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	Email    string `json:"email"`
	Token    string `json:"token"`
	Username string `json:"username"`
	Bio      string `json:"bio"`
	Image    string `json:"image"`
}

type UserController struct {
	svc *service.Service
}

func NewUserController(s *service.Service) *UserController {
	return &UserController{
		svc: s,
	}
}

func (c *UserController) GetUsers(ctx echo.Context) error {
	// users := &[]models.User{}
	// if err := c.repo.Find(ctx.Request().Context(), users); err != nil {
	// 	return ctx.JSON(http.StatusBadRequest, err.Error())
	// }

	return ctx.JSON(http.StatusOK, "users")
}

func (c *UserController) RegisterUser(ctx echo.Context) error {
	var req RegisterUserRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Errors: dto.ErrorBody{Body: []string{err.Error()}},
		})
	}

	user := models.User{
		Username: req.User.Username,
		Email:    req.User.Email,
		Password: req.User.Password,
	}
	err := c.svc.User.RegisterUser(ctx.Request().Context(), user)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, dto.ErrorResponse{
			Errors: dto.ErrorBody{Body: []string{err.Error()}},
		})
	}

	token, err := service.CreateToken(user.Username, user.ID)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, dto.ErrorResponse{
			Errors: dto.ErrorBody{Body: []string{err.Error()}},
		})
	}

	return ctx.JSON(http.StatusCreated, UserResponse{
		Email:    user.Email,
		Username: user.Username,
		Token:    token,
	})
}

func (c *UserController) Login(ctx echo.Context) error {
	var req LoginUserRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, dto.ErrorResponse{
			Errors: dto.ErrorBody{Body: []string{err.Error()}},
		})
	}

	user, err := c.svc.User.LoginUser(ctx.Request().Context(), req.User.Email, req.User.Password)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, dto.ErrorResponse{
			Errors: dto.ErrorBody{Body: []string{err.Error()}},
		})
	}

	token, err := service.CreateToken(user.Username, user.ID)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, dto.ErrorResponse{
			Errors: dto.ErrorBody{Body: []string{err.Error()}},
		})
	}

	return ctx.JSON(http.StatusCreated, UserResponse{
		Email:    user.Email,
		Username: user.Username,
		Token:    token,
	})
}
