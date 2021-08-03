package entities

import "github.com/jinzhu/gorm"

type Transactions struct {
	gorm.Model
	Amount uint `json:"amount"`
	Currency string `json:"currency"`
	SourceId string `json:"source_id"`
	DestinationId string `json:"destination_id"`
	Type string `json:"type"`
	Status     string `json:"status"`
	PaymentId uint
	Payments Payments
}

func (b *Payments) TableName() string {
	return "payments"
}

