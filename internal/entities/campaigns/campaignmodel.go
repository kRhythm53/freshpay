package campaigns

type Campaigns struct {
	Id string                `json:"id"`
	Type string              `json:"type"`
	StartTime int            `json:"start_time"`
	EndTime  int             `json:"end_time"`
	Count int                `json:"count"`
	TransactionNumber int    `json:"transaction_number"`
	IsActive bool            `json:"is_active"`
	MaxCashback int          `json:"max_cashback"`
	PercentageRate int       `json:"percentage_rate"`
	CreatedAt int            `json:"created_at"`
	UpdatedAt int            `json:"updated_at"`
}
func (b *Campaigns) TableName() string{
	return "campaigns"
}