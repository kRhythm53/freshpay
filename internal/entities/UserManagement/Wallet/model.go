package Wallet

type Detail struct {
	EntityId string `json:"entity_id"`
	UserId string `json:"user_id"`
	Balance int `json:"balance"`
	Currency int64 `json:"currency"`
}

const (
	TableName="wallet"
	EntityName="wallet"
)
func(sd *Detail) TableName() string{
	return TableName
}
