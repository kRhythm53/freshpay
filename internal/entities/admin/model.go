package admin

import "gorm.io/gorm"

type Detail struct {
	gorm.Model
	ID          string `gorm:"type:varchar(20)"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
	Email       string `json:"email"`
}

const (
	TableName="admin"
	EntityName="admin"
)
func(sd *Detail) TableName() string{
	return TableName
}