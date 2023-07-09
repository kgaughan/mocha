package models

type Category struct {
	ID    uint `gorm:"primaryKey"`
	User  User
	Title string
}
