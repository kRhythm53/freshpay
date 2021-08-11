package admin

import "gorm.io/gorm"

type Detail struct {
	gorm.Model
	ID          string `gorm:"type:varchar(20)"`
	Name        string
	PhoneNumber string
	Password    string
	Email       string
	IsVerified bool `gorm:"default:false"`
}

const (
	TableName="admin"
	EntityName="admin"
	Prefix="ad"
	IDLengthExcludingPrefix=14
)
func(sd *Detail) TableName() string{
	return TableName
}