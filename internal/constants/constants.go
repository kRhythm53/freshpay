package constants

const (
	PaymentPrefix             = "paymt"
	WalletPrefix              = "wal"
	BankPrefix                = "bnk"
	BeneficiaryPrefix         = "ben"
	TransactionPrefix         = "trans"
	PaymentTypeWalletTransfer = "wallet transfer"
	PaymentTypeBankWithdrawal = "bank withdrawal"
	PaymentTypeAddToWallet    = "add to wallet"
	RazorpayName              = "Razorpay Central Account"
	RazorpayPassword          = "Razorpay123"
	RazorpayPhoneNumber       = "1234567890"
	RazorpayBalance           = 10000000000
	IDLength                  = 14
)

var (
	RzpWalletID string
)

type Error struct {
	Code        string
	Description string
	Source      string
	Reason      string
	Step        string
	Metadata    string
}
type Failure struct {
	Error Error
}
