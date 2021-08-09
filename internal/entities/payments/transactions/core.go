package transactions

import (
	"fmt"
	"github.com/freshpay/internal/constants"
	"github.com/freshpay/internal/entities/payments/payments"
	"github.com/freshpay/internal/entities/payments/utilities"
	"time"
)
func InitiateTransaction(){
	for {
		select {
		case payment:=<-payments.InputPaymentsChannel:
			var err,err2 error
			if payment.Type=="Cashback" || payment.Type=="Refund"{
				err=AddTransactions(payment,"from razorpay account")
			}else{
				err = AddTransactions(payment, "to razorpay account")
				err2 = AddTransactions(payment,"from razorpay account")
			}
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
	transaction.ID= utilities.RandomString(14,constants.TransactionPrefix)
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
	return AddTransactionToDB(transaction)
}

