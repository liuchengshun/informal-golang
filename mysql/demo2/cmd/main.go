package main

import (
	"fmt"
	"log"

	"github.com/liuchengshun/imformal-form/mysql/demo2/model"
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

func InitDB() (*gorm.DB, error) {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4", username, password, host, port, dbname)
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("init database failed, %v", err)
	}

	if err := DB.AutoMigrate(&model.User{}, &model.Company{}, &model.Client{}, &model.CreditCard{}, &model.Email{}, &model.Toy{}); err != nil {
		return nil, fmt.Errorf("auto migrate failed, %v", err)
	}

	return DB, nil
}

func main() {
	DB, err := InitDB()
	if err != nil {
		log.Println("init database failed, got an error: ", err)
		return
	}

	// model.DealUser(DB)

	if err := model.DealClient(DB); err != nil {
		log.Println("deal client failed, ", err)
		return
	}
}
