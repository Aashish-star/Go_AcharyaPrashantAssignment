package main

import (
	"github.com/Aashish-star/Go_AcharyaPrashantAssignment/database"
	"github.com/Aashish-star/Go_AcharyaPrashantAssignment/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDatabase()
	r := gin.Default()
	r.POST("/user/sign-up", handler.SignUp)
	r.POST("/user/sign-in", handler.SignIn)
	r.POST("/user/authorize-token", handler.AuthorizeToken)
	r.POST("/user/refresh-token", handler.RefreshJwtToken)
	r.Run(":8080")

	//for my test
	// fmt.Println("Multiple of two numbers ", utils.Multiple(3, 5))
	// fmt.Println("Factorial of no ", utils.Factorial(5))
	// str := []string{"One", "Two", "Three"}
	// utils.IterateStringArray(str)
	// str = append(str, "Test")
	// str1 := make([]string, 3)
	// str1[0] = "Test"
	// str1[1] = "Test2"

}
