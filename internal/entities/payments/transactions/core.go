package transactions

import (
	"fmt"
	"github.com/freshpay/internal/config"
	"github.com/freshpay/internal/entities/payments/payments"
	"github.com/freshpay/internal/entities/payments/utilities"
	"time"
)
func InitiateTransaction(){
	for {
		select {
		case payment:=<-payments.InputPaymentsChannel:
			err := AddTransactions(payment, "to razorpay account")
			err2 := AddTransactions(payment,"from razorpay account")
			if err != nil || err2!=nil{
				payment.Status="failed"
			}else{
				payment.Status="processed"
			}
			payments.ResultsPaymentsChannel<-payment
		}
	}
}

func AddTransactions(payment *payments.Payments,direction string) (err error) {
	var transaction Transactions
	transaction.ID="trans_"+ utilities.RandomString(14)
	transaction.Amount=payment.Amount
	transaction.Currency=payment.Currency
	if direction=="to razorpay account"{
		transaction.SourceId=payment.SourceId
		transaction.DestinationId="rzp_1234567890abcd"
	}else{
		transaction.SourceId="rzp_1234567890abcd"
		transaction.DestinationId=payment.DestinationId
	}
	transaction.Type=payment.Type
	transaction.Status="processed"
	transaction.PaymentsId=payment.ID
	transaction.CreatedAt=time.Now().Unix()
	transaction.UpdatedAt=time.Now().Unix()
	fmt.Println("transaction : ",transaction)
	fmt.Println("payment:",*payment)
	if err = config.DB.Table("transactions").Create(transaction).Error; err != nil {
		return err
	}
	return nil
}

