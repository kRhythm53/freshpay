package transactions

import (
	"github.com/freshpay/internal/entities/payments/payments"
	"github.com/freshpay/internal/entities/user_management/wallet"
	"github.com/freshpay/utilities"
	"strings"
	"time"
)

func InitiateTransaction() {
	for {
		select {
		case payment := <-payments.InputPaymentsChannel:
			var err error
			if payment.Type == payments.PaymentTypeCashback || payment.Type == payments.PaymentTypeRefund {
				if err = AddTransactions(payment, payments.RzpWalletID, payment.DestinationId);err != nil {
					payment.Status = payments.PaymentStatusFailed
				} else {
					payment.Status = payments.PaymentStatusProcessed
				}
			} else if payment.Type == payments.PaymentTypeAddToWallet || payment.Type == payments.PaymentTypeWalletTransfer || payment.Type == payments.PaymentTypeBankWithdrawal {
				if err = AddTransactions(payment, payment.SourceId, payments.RzpWalletID); err != nil {
					payment.Status = payments.PaymentStatusFailed
				} else if err = AddTransactions(payment, payments.RzpWalletID, payment.DestinationId); err != nil {
					payment.Status = payments.PaymentStatusFailed
				} else {
					payment.Status = payments.PaymentStatusProcessed
				}
			}
			payments.ResultsPaymentsChannel <- payment
		}
	}
}

func AddTransactions(payment *payments.Payments, sourceID string, destinationID string) (err error) {
	var transaction Transactions
	transaction.ID = utilities.CreateID(Prefix, IDLength)
	transaction.Amount = payment.Amount
	transaction.Currency = payment.Currency

	transaction.SourceId = sourceID
	transaction.DestinationId = destinationID

	transaction.Type = payment.Type
	transaction.Status = payments.PaymentStatusProcessed
	transaction.PaymentsId = payment.ID
	transaction.CreatedAt = time.Now().Unix()
	transaction.UpdatedAt = time.Now().Unix()
	if strings.HasPrefix(transaction.SourceId, wallet.Prefix) {
		wallet.UpdateWalletBalance(transaction.SourceId, -1*transaction.Amount)
	}
	if strings.HasPrefix(transaction.DestinationId, wallet.Prefix) {
		wallet.UpdateWalletBalance(transaction.DestinationId, transaction.Amount)
	}
	return AddTransactionToDB(transaction)
}

