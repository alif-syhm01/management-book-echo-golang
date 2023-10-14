package models

type Genre struct {
	ID   uint   `gorm:"primary_key" json:"id" form:"id"`
	Name string `gorm:"not null" json:"name" form:"name" validate:"required"`
}
