package complaints

import (
	"github.com/freshpay/internal/base"
	_ "gorm.io/gorm"
)

type Complaint struct {
	base.Model
	ComplaintType string `json:"complaint_type"`
	Status        string `json:"status"`
	Remark        string `json:"remark"`
	PaymentsId    string `json:"payments_id"`
	UserId        string `json:"user_id"`
	RefundId      string `json:"refund_id"`
	AdminId		  string `json:"admin_id"`
}

func (c *Complaint) TableName() string {
	return "complaint"
}
