package payments

import "gorm.io/gorm"

type Payments struct {
	gorm.Model
	ID            string `gorm:"type:varchar(20)"`
	Amount        uint
	Currency      string
	SourceId      string
	DestinationId string
	Type          string
	Status        string
}


func (b *Payments) TableName() string {
	return "payments"
}