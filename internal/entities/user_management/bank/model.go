package bank

import "gorm.io/gorm"

type Detail struct {
	gorm.Model
	ID            string `gorm:"type:varchar(20)"`
	UserId        string
	BankName      string
	AccountNumber string
	IFSCCode      string
}

const (
	TableName               ="bank"
	EntityName              ="bank"
	Prefix                  ="bnk"
	IDLengthExcludingPrefix =14
)
func(sd *Detail) TableName() string{
	return TableName
}