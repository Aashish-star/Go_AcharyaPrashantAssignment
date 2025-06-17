package handler

import (
	"fmt"

	"github.com/Aashish-star/Go_AcharyaPrashantAssignment/database"
	"github.com/Aashish-star/Go_AcharyaPrashantAssignment/model"
)

func GetUserCountByEmail(Email string) int64 {
	var count int64
	database.DB.Model(&model.UserEntity{}).Where("email = ?", Email).Count(&count)
	fmt.Printf("Total record ", count)
	return count
}

func SaveUser(user *model.UserEntity) int64 {
	var count int64
	database.DB.Where("email= ?", user.Email).Find(&user).Count(&count)
	if count > 0 {
		return 201
	}
	database.DB.Create(&user)
	return 200
}
