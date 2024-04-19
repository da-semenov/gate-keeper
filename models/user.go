package models

type User struct {
	ID      uint   `gorm:"primary_key"`
	Name    string `gorm:"type:varchar(255);not null"`
	Surname string `gorm:"type:varchar(255);not null"`
	Email   string `gorm:"uniqueIndex;not null"`
}
