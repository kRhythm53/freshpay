package payments

import (
	"github.com/freshpay/internal/base"
)

type Payments struct {
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
}

func (b *Payments) TableName() string {
	return "payments"
}
