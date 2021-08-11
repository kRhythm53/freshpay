package transactions

import (
	"github.com/freshpay/internal/constants"
	"github.com/freshpay/internal/entities/payments/payments"
	"github.com/freshpay/internal/entities/user_management/bank"
	"github.com/freshpay/internal/entities/user_management/wallet"
	"github.com/freshpay/utilities"
	"strings"
	"time"
)
func InitiateTransaction(){
	for {
		select {
		case payment:=<-payments.InputPaymentsChannel:
			var err error
			if payment.Type=="Cashback" || payment.Type=="Refund"{
				if strings.HasPrefix(payment.DestinationId, constants.BankPrefix){
					var Bank bank.Detail
					var Wallet wallet.Detail
					err = bank.GetBankById(&Bank, payment.DestinationId)
					if err != nil {
						return
					}
					err = wallet.GetWalletByUserId(&Wallet, Bank.UserId)
					if err != nil {
						return
					}
					payment.DestinationId=Wallet.ID
				}
				err=AddTransactions(payment,constants.RzpWalletID,payment.DestinationId)
				if err != nil {
					payment.Status="failed"
				}else{
					payment.Status="processed"
				}
			}else{
				if err = AddTransactions(payment, payment.SourceId,constants.RzpWalletID);err !=nil{
					payment.Status="failed"
				}else if err = AddTransactions(payment,constants.RzpWalletID,payment.DestinationId);err !=nil{
					payment.Status="failed"
				}else{
					payment.Status="processed"
				}
			}
			payments.ResultsPaymentsChannel<-payment
		}
	}
}

func AddTransactions(payment *payments.Payments,sourceID string, destinationID string) (err error) {
	var transaction Transactions
	transaction.ID= utilities.RandomString(constants.IDLength,constants.TransactionPrefix)
	transaction.Amount=payment.Amount
	transaction.Currency=payment.Currency

	transaction.SourceId=sourceID
	transaction.DestinationId=destinationID

	transaction.Type=payment.Type
	transaction.Status="processed"
	transaction.PaymentsId=payment.ID
	transaction.CreatedAt=time.Now().Unix()
	transaction.UpdatedAt=time.Now().Unix()
	if strings.HasPrefix(transaction.SourceId, constants.WalletPrefix){
		wallet.UpdateWalletBalance(transaction.SourceId,-1*transaction.Amount)
	}
	if strings.HasPrefix(transaction.DestinationId, constants.WalletPrefix){
		wallet.UpdateWalletBalance(transaction.DestinationId,transaction.Amount)
	}
	return AddTransactionToDB(transaction)
}

