package beneficiary

import "gorm.io/gorm"

type Detail struct {
	gorm.Model
	ID                string `gorm:"type:varchar(20)"`
	UserId            string
	BankName          string
	AccountNumber     string
	IFSCCode          string
	AccountHolderName string
}

const (
	TableName="beneficiary"
	EntityName="beneficiary"
)
func(sd *Detail) TableName() string{
	return TableName
}