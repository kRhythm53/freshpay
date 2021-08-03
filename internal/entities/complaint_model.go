package entities

type Complaint struct {
	ComplaintId 		string `gorm:"primaryKey"  json:"id"`
	ComplaintType 				string `json:"complaint_type"`
	PaymentId 			string `json:"payment_id"`
	Status 			string `json:"status"`
	Remark 				string `json:"remark"`
}
func (c *Complaint) TableName() string {
	return "Complaint"
}