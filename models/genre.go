package models

type Genre struct {
	ID   uint   `gorm:"primary_key" json:"id" form:"name"`
	Name string `gorm:"not null" json:"name" form:"name" validate:"required"`
}
