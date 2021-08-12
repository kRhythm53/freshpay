package payments

import (
	"github.com/freshpay/internal/base"
)

type Payments struct {
	base.Model
	Amount        int64  `json:"amount"`
	Currency      string `json:"currency"`
	SourceId      string `json:"source_id"`
	DestinationId string `json:"destination_id"`
	Type          string `json:"type"`
	Status        string `json:"status"`
}

const (
	Prefix                    = "paymt"
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

func (b *Payments) TableName() string {
	return "payments"
}
