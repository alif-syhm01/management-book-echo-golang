package models

type User struct {
	Id       uint   `gorm:"primaryKey" json:"id" form:"id"`
	Name     string `gorm:"not null" json:"name" form:"name"`
	Email    string `gorm:"not null" json:"email" form:"email"`
	Password string `gorm:"not null" json:"password" form:"password"`
}

type UserLogin struct {
	Email    string `gorm:"not null" json:"email" form:"email" validate:"required"`
	Password string `gorm:"not null" json:"password" form:"password" validate:"required"`
}

type UserRegisterResponse struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserLoginResponse struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}
