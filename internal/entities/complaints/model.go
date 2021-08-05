package complaints

import (
	payments2 "github.com/freshpay/internal/entities/payments/payments"
	_ "gorm.io/gorm"
)
type Complaint struct {
	//gorm.Model
	ID            string `gorm:"type:varchar(20)"`
	ComplaintType string
	Status        string
	Remark        string
	PaymentsId    string
	Payments      payments2.Payments
}
func (c *Complaint) TableName() string {
	return "complaint"
}
