package model

type UserEntity struct {
	Id       uint   `gorm:"primaryKey`
	Name     string `json: "name"`
	Email    string `json: "email"`
	Password string `json: "password"`
}
