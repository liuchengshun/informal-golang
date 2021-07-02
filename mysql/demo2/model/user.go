package model

import "gorm.io/gorm"

type User struct {
	ID        string  `json:"id"`
	Name      string  `json:"user"`
	CompanyID int     `gorm:"not null"`
	Company   Company `json:"company" gorm:"foreignkey:CategoryId"`
}

func (u *User) Create(DB *gorm.DB) error {
	return DB.Create(u).Error
}

func (u *User) GetUser(DB *gorm.DB, id string) error {
	return DB.Where("id = ?", id).Preload("")
}
