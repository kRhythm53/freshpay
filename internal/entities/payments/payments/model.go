package payments

var IDPrefix = "paymt_"
var WalletPrefix = "wallt_"
var BankPrefix = "bank_"
var PaymentTypeWalletTransfer = "wallet transfer"
var PaymentTypeBankWithdrawal = "bank withdrawal"
var PaymentTypeAddToWallet = "add to wallet"

type Payments struct {
	//gorm.Model
	ID            string `gorm:"type:varchar(20)"`
	CreatedAt     int64
	UpdatedAt     int64
	Amount        int64
	Currency      string
	SourceId      string `json:"source_id"`
	DestinationId string `json:"destination_id"`
	Type          string
	Status        string
}

func (b *Payments) TableName() string {
	return "payments"
}
