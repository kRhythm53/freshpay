package entities

type Complaint struct {
	gorm.Model
	ComplaintType 		string `json:"complaint_type"`
	Status 			    string `json:"status"`
	Remark 				string `json:"remark"`
	PaymentId 			string `json:"payment_id"`
	Payments Payments
}
func (c *Complaint) TableName() string {
	return "Complaint"
}
