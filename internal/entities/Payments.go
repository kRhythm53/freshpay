package entities

import "github.com/jinzhu/gorm"

type Payments struct {
	gorm.Model
	Id string `json:"id",gorm:"primaryKey"`
	Amount uint `json:"amount"`
	Currency string `json:"currency"`
	SourceId string `json:"source_id"`
	DestinationId string `json:"destination_id"`
	Type string `json:"type"`
	Status     string `json:"status"`
}

func (b *Payments) TableName() string {
	return "payments"
}