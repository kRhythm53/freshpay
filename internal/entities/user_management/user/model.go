package user

import "gorm.io/gorm"

type Detail struct {
	gorm.Model
	ID          string `gorm:"type:varchar(20)"`
	Name        string
	PhoneNumber string
	Password    string
	Email       string
}

const (
	TableName="user"
	EntityName="user"
)
func(sd *Detail) TableName() string{
	return TableName
}