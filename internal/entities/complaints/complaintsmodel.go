package complaints
type Complaints struct{
	Id string           `json:"id"`
	Type string         `json:"type"`
	PaymentId string    `json:"payment_id"`
	Status string `json:"status"`
	Remark string `json:"remark"`
}
func (b *Complaints) TableName() string{
	return "complaints"
}
