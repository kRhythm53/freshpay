package payments

import (
	"github.com/freshpay/internal/entities/user_management/wallet"
	"github.com/freshpay/utilities"
	"testing"
)

func TestGetPaymentType(t *testing.T) {
	var payment Payments
	payment.SourceId = utilities.CreateID(wallet.Prefix, IDLength)
	payment.DestinationId = utilities.CreateID(wallet.Prefix, IDLength)
	if GetPaymentType(&payment) != PaymentTypeWalletTransfer {
		t.Errorf("GetPaymentType(%q)=%q, want %q", &payment, GetPaymentType(&payment), PaymentTypeAddToWallet)
	}
}

func TestGetPaymentsByTime(t *testing.T) {
	var payment []Payments
	if err := GetPaymentsByTime(&payment, "", "", "credit", RzpWalletID); err != nil {
		t.Errorf("GetPaymentsByTime failed")
	}
}

func TestAddPayments(t *testing.T) {
	var payment Payments
	payment.SourceId = utilities.CreateID(wallet.Prefix, IDLength)
	payment.DestinationId = utilities.CreateID(wallet.Prefix, IDLength)
	payment.Amount = 100
	if err := AddPayments(&payment, RzpWalletID); err != nil {
		t.Errorf("AddPayment test failed")
	}
}
