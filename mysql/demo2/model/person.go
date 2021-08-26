package model

type Person struct {
	ID    string  `json:"id" gorm:"primary_key"`
	Name  string  `json:"name"`
	Email []Email `json:"email"`
}
