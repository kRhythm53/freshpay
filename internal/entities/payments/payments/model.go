package payments

import (
	"github.com/freshpay/internal/base"
)

type Payments struct {
	base.Model
	Amount        int64
	Currency      string
	SourceId      string `json:"source_id"`
	DestinationId string `json:"destination_id"`
	Type          string
	Status        string
}

func (b *Payments) TableName() string {
	return "payments"
}
