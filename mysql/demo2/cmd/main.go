package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	username = "lcs_test"
	password = "lcs_test"
	host     = "127.0.0.1"
	port     = "3306"
	dbname   = "test1"
)

type Role struct {
	ID   int    `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
}

type Company struct {
	ID      int    `json:"id" gorm:"primary_key"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type User struct {
	ID        string  `json:"id" gorm:"autoIncrement"`
	Name      string  `json:"user"`
	CompanyID int     `gorm:"not null"`
	Company   Company `json:"company" gorm:"foreignkey:CompanyID"`
}

func InitDB() (*gorm.DB, error) {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4", username, password, host, port, dbname)
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("init database failed, %v", err)
	}

	if err := DB.AutoMigrate(&User{}, &Company{}); err != nil {
		return nil, fmt.Errorf("auto migrate failed, %v", err)
	}

	return DB, nil
}

func (u *User) Create(DB *gorm.DB) error {
	return DB.Create(u).Error
}

func (u *User) GetUsers(DB *gorm.DB) (*[]User, error) {
	users := []User{}
	err := DB.Preload("Company").Find(&users).Error
	return &users, err
}

func GetUser(DB *gorm.DB, userID string) (*User, error) {
	user := &User{}
	err := DB.Joins("Company").First(user, userID).Error
	return user, err
}

func main() {
	DB, err := InitDB()
	if err != nil {
		log.Println("init database failed, got an error: ", err)
		return
	}

	user := User{
		ID:        "1",
		Name:      "ZhangSan",
		CompanyID: 1,
		Company: Company{
			ID:      1,
			Name:    "江南皮革厂",
			Address: "非洲",
		},
	}
	// if err := user.Create(model.DB); err != nil {
	// 	log.Println("create user failed, got an error: ", err)
	// 	return
	// }
	u, err := GetUser(DB, user.ID)
	if err != nil {
		log.Println("get user failed :", err)
		return
	}
	fmt.Printf("u: %v\n", u)

	users, err := user.GetUsers(DB)
	if err != nil {
		log.Println("get users failed, got an error: ", err)
		return
	}
	fmt.Printf("users: %v\n", users)
}
