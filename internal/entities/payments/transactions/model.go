package transactions

import (
	"github.com/freshpay/internal/base"
	"github.com/freshpay/internal/entities/payments/payments"
)

type Transactions struct {
	base.Model
	Amount        int64
	Currency      string
	SourceId      string `json:"source_id"`
	DestinationId string `json:"destination_id"`
	Type          string
	Status        string
	//PaymentsId    string `gorm:"type:varchar(20)"`
	PaymentsId    string `gorm:"type:varchar(20),constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Payments      payments.Payments
}

const (
	Prefix                    = "trans"
	IDLength                  = 14
)


func (b *Transactions) TableName() string {
	return "transactions"
}
