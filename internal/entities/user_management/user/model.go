package user

import (
	"github.com/freshpay/internal/base"
)

type Detail struct {
	base.Model
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