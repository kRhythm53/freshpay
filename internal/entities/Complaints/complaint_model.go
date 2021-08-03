package Complaints

import (
	"github.com/jinzhu/gorm"
	"time"
)
type Complaint struct {
	ID	 string `gorm:"primaryKey"  json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time ``json:"updated_at"
	ComplaintType 		string `json:"complaint_type"`
	Status 			    string `json:"status"`
	Remark 				string `json:"remark"`
	PaymentId 			string `json:"payment_id"`
	Payments Payments
}
func (c *Complaint) TableName() string {
	return "Complaint"
}
