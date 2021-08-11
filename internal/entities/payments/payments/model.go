package payments

import (
	"github.com/freshpay/internal/base"
)

type Payments struct {
	base.Model
	Amount        int64  `json:"amount"`
	Currency      string `json:"currency"`
	SourceId      string `json:"source_id"`
	DestinationId string `json:"destination_id"`
	Type          string `json:"type"`
	Status        string `json:"status"`
}

func (b *Payments) TableName() string {
	return "payments"
}
