package campaigns

var CampaignPrefix = "campn"

type Campaign struct {
	//gorm.Model
	ID                string `gorm:"type:varchar(20)"`
	CampaignType      string `json:"campaign_type"`
	StartTime         int64  `json:"start_time"`
	EndTime           int64  `json:"end_time"`
	Count             int64  `json:"count"`
	TransactionNumber int64 `json:"transaction_number"`
	IsActive          bool   `json:"is_active"`
	MaxCashback       int64  `json:"max_cashback"`
	PercentageRate    int64  `json:"percentage_rate"`
}

func (c *Campaign) TableName() string {
	return "campaign"
}
