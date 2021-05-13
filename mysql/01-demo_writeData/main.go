package main

import (
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID        string `gorm:"primaryKey"`
	Name      string
	Email     string
	Age       uint8
	CreatedAt time.Time
	UpdatedAt time.Time
}

var db *gorm.DB

func main() {
	DBinit()

	user1 := User{
		Name:  "ZhangSan",
		Age:   25,
		Email: "123@fox.com",
	}

	if err := user1.createUser(); err != nil {
		log.Fatal("create user failed, get an error: ", err)
		return
	}

	if users, err := getUsers(); err != nil {
		log.Fatal("get users failed : ", err)
		return
	} else {
		fmt.Println("users:", users)
	}
}

func DBinit() {
	var err error
	dsn := "lcs_test:lcs_test@tcp(127.0.0.1:3306)/lcs?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("open sql error:", err)
	}

	if err = db.AutoMigrate(&User{}); err != nil {
		log.Fatal("table migrate failed")
	}
}

func (u *User) createUser() error {
	u.ID = uuid.NewString()
	return db.Create(&u).Error
}

func getUsers() ([]User, error) {
	users := make([]User, 0)
	err := db.Find(&users).Error
	return users, err
}
