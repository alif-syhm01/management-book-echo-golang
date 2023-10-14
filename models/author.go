package models

type Author struct {
	ID          uint   `gorm:"primaryKey" json:"id" form:"id"`
	Name        string `gorm:"not null" json:"name" form:"name" validate:"required"`
	Nationality string `gorm:"not null" json:"nationality" form:"nationality" validate:"required"`
	Biography   string `gorm:"not null" json:"biography" form:"biography" validate:"required"`
	Book        []Book `json:"books"`
}

type AuthorResponse struct {
	ID          uint   `json:"id" form:"id"`
	Name        string `json:"name" form:"name"`
	Nationality string `json:"nationality" form:"nationality"`
	Biography   string `json:"biography" form:"biography"`
}

type AuthorBookResponse struct {
	ID          uint           `gorm:"primaryKey" json:"id" form:"id"`
	Name        string         `gorm:"not null" json:"name" form:"name" validate:"required"`
	Nationality string         `gorm:"not null" json:"nationality" form:"nationality" validate:"required"`
	Biography   string         `gorm:"not null" json:"biography" form:"biography" validate:"required"`
	Book        []BookResponse `json:"books"`
}
