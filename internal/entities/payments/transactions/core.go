package transactions

import (
	"github.com/freshpay/internal/config"
	"github.com/freshpay/internal/entities/payments/payments"
	"github.com/freshpay/internal/entities/payments/utilities"
)

func AddTransactions(payment payments.Payments,direction string) (err error) {
	var transaction Transactions
	transaction.ID="trans_"+ utilities.RandomString(14)
	transaction.Amount=payment.Amount
	transaction.Currency=payment.Currency
	if direction=="to razorpay account"{
		transaction.SourceId=payment.SourceId
		transaction.DestinationId="rzp account"
	}else{
		transaction.SourceId="rzp account"
		transaction.DestinationId=payment.DestinationId
	}
	transaction.Type=payment.Type
	transaction.Status="processed"
	transaction.PaymentsId=payment.ID

	if err = config.DB.Table("transactions").Create(transaction).Error; err != nil {
		return err
	}
	return nil
}