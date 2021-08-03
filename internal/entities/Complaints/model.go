package Complaints

import (
	"github.com/kshitij-nawandar9/freshpay/internal/entities/Payments"
	_ "gorm.io/gorm"
	"time"
)
type Complaint struct {
	ID	 string `json:"id",gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ComplaintType 		string `json:"complaint_type"`
	Status 			    string `json:"status"`
	Remark 				string `json:"remark"`
	PaymentsId 			string `json:"payment_id"`
	Payments Payments.Payments
}
func (c *Complaint) TableName() string {
	return "complaint"
}
