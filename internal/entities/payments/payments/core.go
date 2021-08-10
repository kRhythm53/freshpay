package payments

import (
	"errors"
	"fmt"
	"github.com/freshpay/internal/config"
	"github.com/freshpay/internal/constants"
	"github.com/freshpay/internal/entities/campaigns"
	"github.com/freshpay/internal/entities/payments/utilities"
	"github.com/freshpay/internal/entities/user_management/bank"
	"github.com/freshpay/internal/entities/user_management/user"
	"github.com/freshpay/internal/entities/user_management/wallet"
	"strconv"
	"strings"
	"time"
)

var InputPaymentsChannel = make(chan *Payments, 1000)
var ResultsPaymentsChannel = make(chan *Payments, 1000)

func AddPayments(payment *Payments) (err error) {
	payment.Type = GetPaymentType(payment)
	payment.Status = "processing"
	payment.ID = utilities.RandomString(14, constants.PaymentPrefix)
	err = ValidityCheck(payment)
	if err != nil {
		return err
	}
	InputPaymentsChannel <- payment

	return AddPaymentToDB(payment)
}

func GetPaymentByID(payment *Payments, id string) (err error) {
	return GetPaymentByIDFromDB(payment, id)
}

func GetPaymentsByTime(payments *[]Payments, from string, to string, TransactionType string, userID string) (err error) {
	var startTime, endTime int64
	if from == "" {
		startTime = 0
	} else {
		startTime, err = strconv.ParseInt(from, 10, 64)
		if err != nil {
			return
		}
	}
	if to == "" {
		endTime = time.Now().Unix()
	} else {
		endTime, err = strconv.ParseInt(to, 10, 64)
	}
	var Wallet wallet.Detail
	wallet.GetWalletByUserId(&Wallet, userID)

	return GetPaymentByTimeFromDB(payments, startTime, endTime, TransactionType, Wallet.ID)
}

func UpdatePayment(payment *Payments) (err error) {
	id, err := GetUserIdFromFundId(payment.SourceId)
	if err != nil {
		return err
	}
	err = UpdateTransactionCount(id)
	if err != nil {
		return err
	}
	if payment.Type != "Cashback" && payment.Type != "Refund" {
		err2 := InitiateCashback(payment)
		if err2 != nil {
			return err2
		}
	}
	return UpdatePaymentToDB(payment)
}

func PaymentReceiver() {
	for {
		select {
		case payment := <-ResultsPaymentsChannel:
			err := UpdatePayment(payment)
			if err != nil {
				return
			}
		}
	}
}

func GetPaymentType(payment *Payments) string {
	if strings.HasPrefix(payment.SourceId, constants.WalletPrefix) {
		if strings.HasPrefix(payment.DestinationId, constants.WalletPrefix) {
			return constants.PaymentTypeWalletTransfer
		} else {
			return constants.PaymentTypeBankWithdrawal
		}
	} else {
		return constants.PaymentTypeAddToWallet
	}
}

func ValidityCheck(payment *Payments) (err error) {
	var balance int
	if strings.HasPrefix(payment.SourceId, constants.WalletPrefix) {
		var Source wallet.Detail
		err := wallet.GetWalletById(&Source, payment.SourceId)
		if err != nil {
			return err
		}
		balance = Source.Balance
		if balance < int(payment.Amount) {
			fmt.Println("low balance")
			return errors.New("low wallet balance")
		}
	} else {
		var Source bank.Detail
		err := bank.GetBankById(&Source, payment.SourceId)
		if err != nil {
			return err
		}
	}
	if strings.HasPrefix(payment.DestinationId, constants.WalletPrefix) {
		var Destination wallet.Detail
		err := wallet.GetWalletById(&Destination, payment.DestinationId)
		if err != nil {
			return err
		}
	} else {
		var Destination bank.Detail
		err := bank.GetBankById(&Destination, payment.DestinationId)
		if err != nil {
			return err
		}
	}

	return nil
}

func InitiateRefund(paymentID string, UserID string) (RefundID string, err error) {
	var RefundPayment Payments
	var payment Payments
	err2 := GetPaymentByID(&payment, paymentID)
	if err2 != nil {
		return "", err2
	}

	var RefundWallet wallet.Detail
	err3 := wallet.GetWalletByUserId(&RefundWallet, UserID)
	if err3 != nil {
		return "", err3
	}

	RefundPayment.ID = utilities.RandomString(14, constants.PaymentPrefix)
	RefundPayment.Amount = payment.Amount
	RefundPayment.Currency = "INR"
	RefundPayment.SourceId = constants.RzpWalletID
	RefundPayment.DestinationId = RefundWallet.ID
	RefundPayment.Type = "Refund"
	RefundPayment.Status = "processing"
	InputPaymentsChannel <- &RefundPayment
	err4:=AddPaymentToDB(&RefundPayment)
	if err4!=nil{
		return "",err4
	}
	return RefundPayment.ID, nil
}

func InitiateCashback(payment *Payments) (err error) {
	var userID string
	if strings.HasPrefix(payment.SourceId, constants.WalletPrefix) {
		var Source wallet.Detail
		err := wallet.GetWalletById(&Source, payment.SourceId)
		if err != nil {
			return err
		}
		userID = Source.UserId
	} else {
		var Source bank.Detail
		err := bank.GetBankById(&Source, payment.SourceId)
		if err != nil {
			return err
		}
		userID = Source.UserId
	}
	Cashback := campaigns.Eligibility(payment.CreatedAt, payment.Amount, userID)
	if Cashback > 0 {
		var CashbackPayment Payments
		CashbackPayment.ID = utilities.RandomString(14, constants.PaymentPrefix)
		CashbackPayment.Amount = int64(Cashback)
		CashbackPayment.Currency = "INR"
		CashbackPayment.SourceId = constants.RzpWalletID
		CashbackPayment.DestinationId = payment.SourceId
		CashbackPayment.Type = "Cashback"
		CashbackPayment.Status = "processing"
		InputPaymentsChannel <- &CashbackPayment
		return AddPaymentToDB(&CashbackPayment)
	}
	return nil
}

func GetUserIdFromFundId(FundId string) (string, error) {
	var userID string
	if strings.HasPrefix(FundId, constants.WalletPrefix) {
		var Source wallet.Detail
		err := wallet.GetWalletById(&Source, FundId)
		if err != nil {
			return "", err
		}
		userID = Source.UserId
	} else {
		var Source bank.Detail
		err := bank.GetBankById(&Source, FundId)
		if err != nil {
			return "", err
		}
		userID = Source.UserId
	}
	return userID, nil
}

func UpdateTransactionCount(userID string) (err error) {
	var User user.Detail
	err = user.GetUserById(&User, userID)
	if err != nil {
		return
	}
	User.NumberOfTransactions = User.NumberOfTransactions + 1
	config.DB.Table("user").Save(&User)
	return nil
}

func CreateRzpAccount()(err error){
	var RZP user.Detail
	RZP.Name="Razorpay Central Account"
	RZP.Password="Razorpay123"
	RZP.PhoneNumber="1234567890"
	err = user.SignUp(&RZP)
	if err != nil {
		return err
	}
	var RZPWallet wallet.Detail
	err = wallet.GetWalletByUserId(&RZPWallet, RZP.ID)
	if err != nil {
		return err
	}
	constants.RzpWalletID=RZPWallet.ID
	wallet.UpdateWalletBalance(constants.RzpWalletID,10000000000)
	return nil
}