package campaigns

import "gorm.io/gorm"

var ComplaintPrefix = "cmplt_"
type Campaign struct {
	gorm.Model
	ID                string    `gorm:"type:varchar(20)"`
	CampaignType      string
	StartTime         int64
	EndTime           int64
	Count             int64
	TransactionNumber string
	IsActive          bool
	MaxCashback       int64
	PercentageRate    int64
}
func (c *Campaign) TableName() string {
	return "campaign"
}