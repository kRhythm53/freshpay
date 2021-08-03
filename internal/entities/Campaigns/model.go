package Campaigns

import (
	_ "gorm.io/gorm"
	"time"
)

type Campaign struct {
	ID                string    `json:"id",gorm:"primaryKey"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	CampaignType      string    `json:"name"`
	StartTime         time.Time `json:"start_time"`
	EndTime           time.Time `json:"end_time"`
	Count             string    `json:"count"`
	TransactionNumber string    `json:"transaction_number"`
	IsActive          string    `json:"is_active"`
	MaxCashback       string    `json:"max_cashback"`
	PercentageRate    string    `json:"percentage_rate"`
}
func (c *Campaign) TableName() string {
	return "Campaign"
}