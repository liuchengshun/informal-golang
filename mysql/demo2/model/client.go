package model

import (
	"fmt"

	"gorm.io/gorm"
)

type Client struct {
	ID         string     `json:"id" gorm:"primary_key;type:varchar(32)"`
	Name       string     `json:"name"`
	Emails     []Email    `json:"email" gorm:"foreignKey:value"`
	CreditCard CreditCard `json:"credit_card"`
}

type Email struct {
	Value string `json:"value" gorm:"primary_key"`
	Name  string `json:"name"`
}

type Toy struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	OwnerID   string `json:"owner_id"`
	OwnerType string `json:"owner_type"`
}

type CreditCard struct {
	ID       string `json:"id"`
	Number   string `json:"number"`
	ClientID string `json:"client_id"`
}

func (c *Client) Create(DB *gorm.DB) error {
	// DB.Model(&c).Association("Emails")
	return DB.Create(c).Error
}

func (c *Client) Update(DB *gorm.DB) error {
	return DB.Save(c).Error
}

func (c *Client) Delete(DB *gorm.DB) error {
	return DB.Where("id = ?", c.ID).Delete(c).Error
}

func (c *Client) GetEmailAssociation(DB *gorm.DB) ([]Email, error) {
	emails := make([]Email, 0)
	err := DB.Model(c).Association("Emails").Find(emails)
	return emails, err
}

func GetClient(DB *gorm.DB, clientID string) (*Client, error) {
	c := &Client{}
	err := DB.Preload("Emails").Preload("CreditCard").First(c, clientID).Error
	return c, err
}

func GetClients(DB *gorm.DB) ([]Client, error) {
	var cs []Client
	err := DB.Preload("Emails").Preload("CreditCard").Find(&cs).Error
	return cs, err
}

func DealClient(DB *gorm.DB) error {
	c := &Client{
		ID:   "013",
		Name: "WangWu",
		Emails: []Email{
			{Value: "abc123@qq.com", Name: "heieihei"},
			{Value: "scd123@qq.com", Name: "heieihei"},
			{Value: "fdsfa@qq.com", Name: "heieihei"},
			{Value: "cxgasrg@qq.com", Name: "heieihei"},
		},
		CreditCard: CreditCard{
			ID:     "000",
			Number: "000",
		},
	}
	// if err := c.Create(DB); err != nil {
	// 	return err
	// }

	// c.Name = "Lisi"
	// c.CreditCard.Number = "002"
	// if err := c.Update(DB); err != nil {
	// 	return err
	// }

	// if err := c.Delete(DB); err != nil {
	// 	return err
	// }

	// emails, err := c.GetEmailAssociation(DB)
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("emails: %v\n", emails)

	clientCase1, err := GetClient(DB, c.ID)
	if err != nil {
		return err
	}
	fmt.Printf("clientCase1: %v\n", clientCase1)

	// clients, err := GetClients(DB)
	// if err != nil {
	// 	return err
	// }
	// fmt.Println("clients: ", clients)

	return nil
}
