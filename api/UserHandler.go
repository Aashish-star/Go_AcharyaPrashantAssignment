package api

import (
	"net/http"

	"github.com/Aashish-star/Go_AcharyaPrashantAssignment/internal/model"
	"github.com/Aashish-star/Go_AcharyaPrashantAssignment/internal/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{UserService: userService}
}

func (uc *UserController) RegisterRoutes(r *gin.Engine) {
	userGroup := r.Group("/api/user")
	{
		userGroup.POST("/sign-up", uc.SignUp)
		userGroup.POST("/sign-in", uc.SignIn)
		userGroup.POST("/authorize-token", uc.AuthorizeToken)
		userGroup.POST("/refresh-token", uc.RefreshToken)
		userGroup.GET("/revoke-token", uc.RevokeToken)
	}
}

func (uc *UserController) SignUp(c *gin.Context) {
	var req model.UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	uc.UserService.CreateUser(c, req)
}

func (uc *UserController) SignIn(c *gin.Context) {
	var req model.UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	uc.UserService.SignIn(c, req)
}

func (uc *UserController) AuthorizeToken(c *gin.Context) {
	var req model.UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token := c.GetHeader("Authorization")
	uc.UserService.AuthorizeToken(c, token, req)
}

func (uc *UserController) RefreshToken(c *gin.Context) {
	var req model.JwtRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	uc.UserService.RefreshToken(c, req.RefreshToken)
}

func (uc *UserController) RevokeToken(c *gin.Context) {
	token := c.GetHeader("Authorization")
	uc.UserService.RevokeToken(c, token)
}
