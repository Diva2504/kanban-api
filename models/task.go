package models

type Task struct {
	gorm.Model
	Title       string `gorm:"not null" validate:"required"`
	Description string `gorm:"not null" validate:"required"`
	Status      bool   `gorm:"not null" validate:"required"`
	UserId      uint
	CategoryId  uint
	User        *User     `gorm:"foreignKey:UserId"`
	Category    *Category `gorm:"foreignKey:CategoryId"`
}
