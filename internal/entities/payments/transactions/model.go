package transactions

import (
	"github.com/freshpay/internal/base"
	"github.com/freshpay/internal/entities/payments/payments"
)

type Transactions struct {
	//gorm.Model
	//ID            string `gorm:"type:varchar(20)"`
	//CreatedAt     int64
	//UpdatedAt     int64
	base.Model
	Amount        int64
	Currency      string
	SourceId      string `json:"source_id"`
	DestinationId string `json:"destination_id"`
	Type          string
	Status        string
	PaymentsId    string `gorm:"type:varchar(20),constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Payments      payments.Payments
}
func (b *Transactions) TableName() string {
	return "transactions"
}
