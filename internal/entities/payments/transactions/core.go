package transactions

import (
	"github.com/freshpay/internal/constants"
	"github.com/freshpay/internal/entities/payments/payments"
	"github.com/freshpay/internal/entities/payments/utilities"
	"github.com/freshpay/internal/entities/user_management/bank"
	"github.com/freshpay/internal/entities/user_management/wallet"
	"strings"
	"time"
)
func InitiateTransaction(){
	for {
		select {
		case payment:=<-payments.InputPaymentsChannel:
			var err,err2 error
			if payment.Type=="Cashback" || payment.Type=="Refund"{
				if strings.HasPrefix(payment.DestinationId, constants.BankPrefix){
					var Bank bank.Detail
					var Wallet wallet.Detail
					err := bank.GetBankById(&Bank, payment.DestinationId)
					if err != nil {
						return
					}
					err = wallet.GetWalletByUserId(&Wallet, Bank.UserId)
					if err != nil {
						return
					}
					payment.DestinationId=Wallet.ID
				}
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
		transaction.DestinationId=constants.RzpWalletID
		if strings.HasPrefix(transaction.SourceId, constants.WalletPrefix){
			wallet.UpdateWalletBalance(transaction.SourceId,-1*transaction.Amount)
		}
	}else{
		transaction.SourceId=constants.RzpWalletID
		transaction.DestinationId=payment.DestinationId
		if strings.HasPrefix(transaction.DestinationId, constants.WalletPrefix){
			//fmt.Println(transaction.DestinationId,transaction.Amount)
			wallet.UpdateWalletBalance(transaction.DestinationId,transaction.Amount)
		}
	}
	transaction.Type=payment.Type
	transaction.Status="processed"
	transaction.PaymentsId=payment.ID
	transaction.CreatedAt=time.Now().Unix()
	transaction.UpdatedAt=time.Now().Unix()
	return AddTransactionToDB(transaction)
}

