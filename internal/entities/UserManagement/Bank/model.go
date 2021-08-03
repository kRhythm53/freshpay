package Bank

type Detail struct {
	EntityId string `json:"entity_id"`
	UserId string `json:"user_id"`
	BankName string `json:"bank_name"`
	AccountNumber string `json:"account_number"`
	IFSCCode string `json:"ifsc_code"`
}

const (
	TableName="bank"
	EntityName="bank"
)
func(sd *Detail) TableName() string{
	return TableName
}