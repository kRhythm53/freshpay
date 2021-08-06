package wallet

import "gorm.io/gorm"

type Detail struct {
	gorm.Model
	ID       string `gorm:"type:varchar(20)"`
	UserId   string
	Balance  int
	Currency string
}

const (
	TableName="wallet"
	EntityName="wallet"
)
func(sd *Detail) TableName() string{
	return TableName
}
