package user

import "gorm.io/gorm"

type Detail struct {
	gorm.Model
	ID          string `gorm:"type:varchar(20)"`
	Name        string
	PhoneNumber string
	Password    string
	Email       string
	NumberOfTransactions int64 `gorm:"default:0"`
}

const (
	TableName="user"
	EntityName="user"
	Prefix="user"
)
func(sd *Detail) TableName() string{
	return TableName
}