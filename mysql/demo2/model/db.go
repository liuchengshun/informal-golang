package model

import (
	"fmt"

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

var DB *gorm.DB

func InitDB() error {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4", username, password, host, port, dbname)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("init database failed, %v", err)
	}

	if err := DB.AutoMigrate(&User{}, &Company{}); err != nil {
		return fmt.Errorf("auto migrate failed, %v", err)
	}

	return nil
}
