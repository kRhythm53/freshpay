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
	PaymentsId    string `gorm:"default:''",json:"payments_id"`
	UserId        string `gorm:"default:''",json:"user_id"`
	RefundId      string `gorm:"default:''",json:"refund_id"`
	AdminId		  string `gorm:"default:''",json:"admin_id"`
}

func (c *Complaint) TableName() string {
	return "complaint"
}
