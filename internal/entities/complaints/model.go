package complaints

import (
	"github.com/freshpay/internal/entities/payments/payments"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
)
type Complaint struct {
	gorm.Model
	ID            string `gorm:"type:varchar(20)"`
	ComplaintType string
	Status        string
	Remark        string
	PaymentsId    string
	Payments      payments.Payments
}
func (c *Complaint) TableName() string {
	return "complaint"
}
