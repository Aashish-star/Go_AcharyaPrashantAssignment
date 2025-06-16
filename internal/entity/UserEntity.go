package emtity

type UserDetail struct {
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Email    string `gorm:"column:email" json:"email"`
	UserName string `gorm:"column:userName" json:"userName"`
	Password string `gorm:"column:password" json:"password"`
}

func (UserDetail) TableName() string {
	return "user_detail"
}
