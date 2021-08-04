package session

import "gorm.io/gorm"

type Detail struct {
	gorm.Model
	ID         string `gorm:"type:varchar(20)"`
	UserId     string
	ExpireTime uint64
}

const (
	TableName="session"
	EntityName="session"
)

func(sd *Detail) TableName() string{
	return TableName
}