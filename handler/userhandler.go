package handler

import (
	"fmt"
	"net/http"

	"github.com/Aashish-star/Go_AcharyaPrashantAssignment/model"
	"github.com/Aashish-star/Go_AcharyaPrashantAssignment/utils"
	"github.com/gin-gonic/gin"
)

const TOKEN_IS_VALID = "Token is valid"

func SignUp(c *gin.Context) {
	var request model.UserRequest
	error := c.ShouldBindJSON(&request)
	fmt.Printf("newUser request ", request)
	if error != nil {
		fmt.Printf("Error found")
		c.JSON(http.StatusBadGateway, "Error in CreateUser")
		return
	}
	//creating entity from input request
	var dbEntity model.UserEntity
	dbEntity.Email = request.Email
	dbEntity.Name = request.Name
	dbEntity.Password = request.Password

	//getting db connection

	var response model.UserResponse
	statusCode := SaveUser(&dbEntity)
	if statusCode != 200 {
		response.IsSucess = false
		response.Message = "User already present"
	} else {
		response.IsSucess = true
		response.Message = "User created successfully"
	}
	c.JSON(http.StatusCreated, response)
}

func SignIn(c *gin.Context) {
	var request model.UserRequest
	error := c.ShouldBindJSON(&request)
	fmt.Printf("newUser request ", request)
	if error != nil {
		fmt.Printf("Error found")
		c.JSON(http.StatusBadGateway, "Invalid request")
		return
	}
	count := GetUserCountByEmail(request.Email)
	var response model.UserResponse
	if count > 0 {
		accessToken := utils.GenerateJWT(request.Name)
		refreshToken := utils.GenerateRefreshToken(request.Name)
		fmt.Println("Generated Access Token ", accessToken)
		response.IsSucess = true
		response.Message = "Successfully generated"
		response.AccessToken = accessToken
		response.RefreshToken = refreshToken
	} else {
		response.IsSucess = false
		response.Message = "Not able to generate token"
	}
	c.JSON(http.StatusOK, response)

}

func RefreshJwtToken(c *gin.Context) {
	token := c.GetHeader("Authorization")
	var request model.UserRequest
	var response model.UserResponse
	c.ShouldBindJSON(&request)
	tokenUsername := utils.GetUsernameFromToken(token)
	fmt.Println("tokenUsername ", tokenUsername)

	message := utils.ValidateToken(token)

	if tokenUsername != request.Name {
		fmt.Println("request ", request)
		response.IsSucess = false
		response.Message = "Invalid username"
	} else if message == TOKEN_IS_VALID {
		accessToken := utils.GenerateJWT(request.Name)
		refreshToken := utils.GenerateRefreshToken(request.Name)
		fmt.Println("Generated Access Token ", accessToken)
		response.IsSucess = true
		response.Message = "Successfully generated"
		response.AccessToken = accessToken
		response.RefreshToken = refreshToken
	} else {
		response.IsSucess = false
	}
	c.JSON(http.StatusOK, response)
}

func AuthorizeToken(c *gin.Context) {
	token := c.GetHeader("Authorization")
	response := utils.ValidateToken(token)
	c.JSON(http.StatusOK, response)
}
