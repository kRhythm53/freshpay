package complaints

import (
	"github.com/freshpay/internal/entities/payments/payments"
	_ "gorm.io/gorm"
)
type Complaint struct {
	ID            string `gorm:"type:varchar(20)"`
	CreatedAt     int64
	UpdatedAt     int64
	ComplaintType string
	Status        string
	Remark        string
	PaymentsId    string
	RefundId	  string
	Payments      payments.Payments
}
func (c *Complaint) TableName() string {
	return "complaint"
}
