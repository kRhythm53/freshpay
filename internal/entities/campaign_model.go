package entities

type Campaign struct {
	gorm.Model
	CampaignType 				string `json:"name"`
	StartTime 			string `json:"start_time"`
	EndTime 			string `json:"end_time"`
	Count 				string `json:"count"`
	TransactionNumber  string `json:"transaction_number"`
	IsActive 			string `json:"is_active"`
	MaxCashback		string `json:"max_cashback"`
	PercentageRate		string `json:"percentage_rate"`
}
func (c *Campaign) TableName() string {
	return "Campaign"
}