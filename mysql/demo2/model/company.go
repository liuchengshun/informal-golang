package model

type Company struct {
	ID      int    `gorm:"primary_key"`
	Address string `json:"address"`
}
