package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/omnlgy/go-hris-payroll-system/internal/dto"
	"github.com/omnlgy/go-hris-payroll-system/internal/service"
)

type AuthController struct {
	authService *service.AuthService
}

func NewAuthController(authService *service.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (c *AuthController) Login(ctx *gin.Context) {
	var body dto.LoginRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(400, dto.BadRequestResponse{
			Message: "Invalid request body",
			Errors: []struct {
				Field   string `json:"field"`
				Message string `json:"message"`
			}{
				{
					Field:   "body",
					Message: err.Error(),
				},
			},
		})
		return
	}

	token, err := c.authService.Login(body.Email, body.Password)
	if err != nil {
		ctx.JSON(401, dto.UnauthorizedResponse{
			Message: "Invalid credentials",
		})
		return
	}

	ctx.JSON(200, dto.SuccessResponse{
		Message: "Login successful",
		Data:    token,
	})
}

func (c *AuthController) Logout(ctx *gin.Context) {

}
