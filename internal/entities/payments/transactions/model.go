package transactions

import (
	payments2 "github.com/freshpay/internal/entities/payments/payments"
	"gorm.io/gorm"
)

type Transactions struct {
	gorm.Model
	ID            string `gorm:"type:varchar(20)"`
	Amount        int64
	Currency      string
	SourceId      string
	DestinationId string
	Type          string
	Status        string
	PaymentsId    string `gorm:"type:varchar(20)"`
	Payments      payments2.Payments
}

func (b *Transactions) TableName() string {
	return "transactions"
}
