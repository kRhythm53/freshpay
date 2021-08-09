package transactions

import (
	"github.com/freshpay/internal/constants"
	"github.com/freshpay/internal/entities/payments/payments"
	"github.com/freshpay/internal/entities/payments/utilities"
	"github.com/freshpay/internal/entities/user_management/wallet"
	"time"
)
func InitiateTransaction(){
	for {
		select {
		case payment:=<-payments.InputPaymentsChannel:
			var err,err2 error
			if payment.Type=="Cashback" || payment.Type=="Refund"{
				//fmt.Println("initiating CB")
				err=AddTransactions(payment,"from razorpay account")
				if err != nil {
					payment.Status="failed"
				}else{
					payment.Status="processed"
				}
			}else{
				err = AddTransactions(payment, "to razorpay account")
				err2 = AddTransactions(payment,"from razorpay account")
				if err != nil || err2!=nil{
					payment.Status="failed"
				}else{
					payment.Status="processed"
				}
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
		transaction.DestinationId="wal_Mh5gqYDWlNBYWq"
		wallet.UpdateWalletBalance(transaction.SourceId,-1*transaction.Amount)
	}else{
		transaction.SourceId="wal_Mh5gqYDWlNBYWq"
		transaction.DestinationId=payment.DestinationId
		wallet.UpdateWalletBalance(transaction.DestinationId,transaction.Amount)
	}
	transaction.Type=payment.Type
	transaction.Status="processed"
	transaction.PaymentsId=payment.ID
	transaction.CreatedAt=time.Now().Unix()
	transaction.UpdatedAt=time.Now().Unix()
	//fmt.Println("transaction : ",transaction)
	//fmt.Println("payment:",*payment)
	return AddTransactionToDB(transaction)
}

