package complaints

import (
	"github.com/freshpay/internal/base"
	_ "gorm.io/gorm"
)
type Complaint struct {
	base.Model
	ComplaintType string
	Status        string
	Remark        string
	PaymentsId    string
	UserId        string
	RefundId      string
}
func (c *Complaint) TableName() string {
	return "complaint"
}
