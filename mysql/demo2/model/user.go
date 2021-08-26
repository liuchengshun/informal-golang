package model

import (
	"fmt"
	"log"

	"gorm.io/gorm"
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

// create
func (u *User) Create(DB *gorm.DB) error {
	return DB.Create(u).Error
}

// update user
func (u *User) Update(DB *gorm.DB) error {
	return DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(u).Error
}

// get user
func GetUsers(DB *gorm.DB) (*[]User, error) {
	users := []User{}
	err := DB.Joins("Company").Find(&users).Error
	return &users, err
}

func GetUser(DB *gorm.DB, userID string) (*User, error) {
	user := &User{}
	err := DB.Joins("Company").First(user, userID).Error
	return user, err
}

// update company
func (c *Company) Update(DB *gorm.DB) error {
	return DB.Save(c).Error
}

// get company
func GetCompany(DB *gorm.DB, id int) (*Company, error) {
	c := &Company{}
	DB.Where("id = ?", id).First(c)
	return c, nil
}

func NewCompany(ID int, Name, Address string) *Company {
	return &Company{
		ID:      ID,
		Name:    Name,
		Address: Address,
	}
}

func NewUser(ID, Name string, company Company) *User {
	return &User{
		ID:        ID,
		Name:      Name,
		CompanyID: company.ID,
		Company:   company,
	}
}

func DealUser(DB *gorm.DB) {
	company_case2 := NewCompany(2, "傻子保护协会", "沙湖")
	user_case2 := NewUser("2", "傻子2号", *company_case2)

	// if err := user_case2.Create(DB); err != nil {
	// 	log.Println("create user2 failed, ", err)
	// 	return
	// }

	// user_case2.Company.Name = "砍树联盟"
	// if err := company_case2.Update(DB); err != nil {
	// 	log.Println("update company failed, ", err)
	// 	return
	// }

	// c, err := GetCompany(DB, user_case2.Company.ID)
	// if err != nil {
	// 	log.Println("get company error", err)
	// 	return
	// }
	// fmt.Printf("c: %v\n", c)

	user_case2.Name = "光头强"
	user_case2.Company.Name = "砍树联盟"
	if err := user_case2.Update(DB); err != nil {
		log.Println("update user failed, ", err)
		return
	}

	users, err := GetUsers(DB)
	if err != nil {
		log.Println("get users failed, ", err)
		return
	}
	fmt.Printf("users: %v\n", users)
}
