package constants

type Model struct {
	ID        string `gorm:"type:varchar(20)"`
	CreatedAt int64
	UpdatedAt int64
	DeletedAt int64
}

const (
	IDPrefix = "paymt"
	WalletPrefix = "wallt"
	BankPrefix = "bank"
	TransactionPrefix="trans"
	PaymentTypeWalletTransfer = "wallet transfer"
	PaymentTypeBankWithdrawal = "bank withdrawal"
	PaymentTypeAddToWallet = "add to wallet"
)
