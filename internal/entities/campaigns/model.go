package campaigns

import (
	"gorm.io/gorm"
	_ "gorm.io/gorm"
	"time"
)

type Campaign struct {
	gorm.Model
	ID                string    `gorm:"type:varchar(20)"`
	CampaignType      string
	StartTime         time.Time
	EndTime           time.Time
	Count             string
	TransactionNumber string
	IsActive          string
	MaxCashback       string
	PercentageRate    string
}
func (c *Campaign) TableName() string {
	return "campaign"
}