package payments

import (
	"errors"
	"fmt"
	"github.com/freshpay/internal/config"
	"github.com/freshpay/internal/entities/campaigns"
	"github.com/freshpay/internal/entities/user_management/bank"
	"github.com/freshpay/internal/entities/user_management/beneficiary"
	"github.com/freshpay/internal/entities/user_management/user"
	"github.com/freshpay/internal/entities/user_management/wallet"
	"github.com/freshpay/utilities"
	"strconv"
	"strings"
	"sync"
	"time"
)

//Flow of payments  :
// Initiate a payment -> validity checks -> push to InputPaymentChannel
// Transaction module receives the payment request from channel and initiates transaction steps
// After successful completion of transactions, transaction module pushes the payment to ResultsPaymentChannel
// PaymentReceiver then receives the updated payment struct and updates the DB
// Check if the payment is eligible for a cashback, if yes a cashback is initiated by simply adding a payment request to InputPaymentChannel
// Admin can initiate by refund by initiating a payment to InputPaymentChannel

var InputPaymentsChannel = make(chan *Payments, 1000)
var ResultsPaymentsChannel = make(chan *Payments, 1000)
var mutex = &sync.Mutex{}

// AddPayments :Initiate payment after validity checks, with payment status as processing and push the payment struct to channel
func AddPayments(payment *Payments, userId string) (err error) {
	payment.Type = GetPaymentType(payment)
	payment.Status = PaymentStatusProcessing
	payment.ID = utilities.CreateID(Prefix, IDLength)
	err = ValidityCheck(payment, userId)
	if err != nil {
		return err
	}
	err = AddPaymentToDB(payment)
	if err != nil {
		return err
	}
	InputPaymentsChannel <- payment
	return nil
}

// GetPaymentByID :Get payment for the given payment id
func GetPaymentByID(payment *Payments, id string) (err error) {
	return GetPaymentByIDFromDB(payment, id)
}

// GetPaymentsByTime :Get all payments for logged in user, based on time and type
// from = {int epoch value}, to = {int, epoch value}, type=credit/debit
// if from is not specified , consider from time 0
// if to is not specified , consider till current time
// if type is not specified , consider both debit and credit transactions
func GetPaymentsByTime(payments *[]Payments, from string, to string, TransactionType string, userID string) (err error) {
	var startTime, endTime int64
	if from == "" {
		startTime = 0
	} else {
		startTime, err = strconv.ParseInt(from, 10, 64)
		if err != nil {
			return errors.New("bad request")
		}
	}
	if to == "" {
		endTime = time.Now().Unix()
	} else {
		endTime, err = strconv.ParseInt(to, 10, 64)
		if err != nil {
			return errors.New("bad request")
		}
	}
	var Wallet wallet.Detail
	err = wallet.GetWalletByUserId(&Wallet, userID)
	if err != nil {
		return errors.New("could not fetch wallet details")
	}

	if TransactionType == "credit" {
		return GetPaymentByTimeCreditFromDB(payments, startTime, endTime, Wallet.ID)
	} else if TransactionType == "debit" {
		return GetPaymentByTimeDebitFromDB(payments, startTime, endTime, Wallet.ID)
	} else {
		return GetPaymentByTimeFromDB(payments, startTime, endTime, Wallet.ID)
	}
}

// PaymentReceiver : A receiver go routine after all transactions have been completed
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

// UpdatePayment :Update payment in DB after all transactions are completed successfully,
// increment transaction count of user and check for eligible cashback offers
func UpdatePayment(payment *Payments) (err error) {
	if err = UpdateTransactionCount(payment); err != nil {
		return errors.New("could not update transaction count")
	}
	if err = UpdatePaymentToDB(payment); err != nil {
		return err
	}
	if payment.Type != PaymentTypeCashback && payment.Type != PaymentTypeRefund {
		if err = InitiateCashback(payment); err != nil {
			return err
		}
	}
	return nil
}

// GetPaymentType :i.e. wallet-to-wallet / cash withdrawal / add to wallet / cashback / refund
func GetPaymentType(payment *Payments) string {
	if strings.HasPrefix(payment.SourceId, wallet.Prefix) {
		if strings.HasPrefix(payment.DestinationId, wallet.Prefix) {
			return PaymentTypeWalletTransfer
		} else {
			return PaymentTypeBankWithdrawal
		}
	} else {
		return PaymentTypeAddToWallet
	}
}

// ValidityCheck :
//1. source wallet/bank does not exist or does not belong to user
//2. destination bank/wallet/beneficiary does not exist
//3. negative amount
//4. transaction amount is greater than wallet balance
func ValidityCheck(payment *Payments, userId string) (err error) {
	var SourceUserId, DestinationUserId string

	if SourceUserId, err = GetUserIdFromFundId(payment.SourceId); err != nil {
		return err
	}
	if SourceUserId != userId {
		return errors.New("source does not belong to user")
	}

	if DestinationUserId, err = GetUserIdFromFundId(payment.DestinationId); err != nil {
		return err
	}
	if DestinationUserId != userId {
		return errors.New("destination does not belong to user")
	}

	if payment.Amount < 0 {
		return errors.New("payment amount invalid")
	}
	if strings.HasPrefix(payment.SourceId, wallet.Prefix) {
		var Source wallet.Detail
		if err = wallet.GetWalletById(&Source, payment.SourceId); err != nil {
			return errors.New("source wallet does not exist")
		}
		balance := int64(Source.Balance)
		if balance < payment.Amount {
			return errors.New("low wallet balance")
		}
	}
	return nil
}

func GetUserIdFromFundId(FundId string) (string, error) {
	var userID string
	if strings.HasPrefix(FundId, wallet.Prefix) {
		var Source wallet.Detail
		if err := wallet.GetWalletById(&Source, FundId); err != nil {
			return "", errors.New("wallet does not exist")
		}
		userID = Source.UserId
	} else if strings.HasPrefix(FundId, bank.Prefix) {
		var Source bank.Detail
		if err := bank.GetBankById(&Source, FundId); err != nil {
			return "", err
		}
		userID = Source.UserId
	} else if strings.HasPrefix(FundId, beneficiary.Prefix) {
		var Source beneficiary.Detail
		if err := beneficiary.GetBeneficiaryById(&Source, FundId); err != nil {
			return "", err
		}
		userID = Source.UserId
	}
	if userID!=""{return userID,nil}
	return userID, errors.New("invalid fund id")
}

// InitiateRefund :A refund is initiated by Admin for given payment id and user id
func InitiateRefund(paymentID string, UserID string) (RefundID string, err error) {
	var RefundPayment Payments
	var payment Payments
	err = GetPaymentByID(&payment, paymentID)
	if err != nil {
		return "", errors.New("failed to get payment details")
	}

	var RefundWallet wallet.Detail
	err = wallet.GetWalletByUserId(&RefundWallet, UserID)
	if err != nil {
		return "", err
	}

	RefundPayment.ID = utilities.CreateID(Prefix, IDLength)
	RefundPayment.Amount = payment.Amount
	RefundPayment.SourceId = RzpWalletID
	RefundPayment.DestinationId = RefundWallet.ID
	RefundPayment.Type = PaymentTypeRefund
	RefundPayment.Status = PaymentStatusProcessing
	err = AddPaymentToDB(&RefundPayment)
	if err != nil {
		return "", err
	}
	InputPaymentsChannel <- &RefundPayment
	return RefundPayment.ID, nil
}

// InitiateCashback :Checks with campaign module if the payments is eligible for a cashback,
// if yes, initiate a cashback to user wallet
func InitiateCashback(payment *Payments) (err error) {
	userID, err := GetUserIdFromFundId(payment.SourceId)

	Cashback := campaigns.Eligibility(payment.CreatedAt, payment.Amount, userID)
	if Cashback > 0 {
		var DestinationId string
		if strings.HasPrefix(payment.SourceId, wallet.Prefix) {
			DestinationId = payment.SourceId
		} else {
			var Wallet wallet.Detail
			if err = wallet.GetWalletByUserId(&Wallet, userID); err != nil {
				return err
			}
			DestinationId = Wallet.ID
		}
		var CashbackPayment Payments
		CashbackPayment.ID = utilities.CreateID(Prefix, IDLength)
		CashbackPayment.Amount = int64(Cashback)
		CashbackPayment.SourceId = RzpWalletID
		CashbackPayment.DestinationId = DestinationId
		CashbackPayment.Type = PaymentTypeCashback
		CashbackPayment.Status = PaymentStatusProcessing
		InputPaymentsChannel <- &CashbackPayment
		return AddPaymentToDB(&CashbackPayment)
	}
	return nil
}

func UpdateTransactionCount(payment *Payments) (err error) {
	id, err := GetUserIdFromFundId(payment.SourceId)
	if err != nil {
		return errors.New("failed to update payment , could not find user id for given fund id")
	}

	var User user.Detail
	if err = user.GetUserById(&User, id); err != nil {
		return errors.New("could not find the user")
	}

	mutex.Lock()
	User.NumberOfTransactions = User.NumberOfTransactions + 1
	mutex.Unlock()

	config.DB.Table("user").Save(&User)
	return nil
}

func CreateRzpAccount() (err error) {
	var RZP user.Detail
	if err = user.GetUserByPhoneNumber(&RZP, RazorpayPhoneNumber); err != nil {
		RZP.Name = RazorpayName
		RZP.Password = RazorpayPassword
		RZP.PhoneNumber = RazorpayPhoneNumber
		RZP.IsVerified = true
		err = user.SignUp(&RZP)
		if err != nil {
			return err
		}
	} else {
		fmt.Println("Razorpay wallet already exists")
	}
	var RZPWallet wallet.Detail
	err = wallet.GetWalletByUserId(&RZPWallet, RZP.ID)
	if err != nil {
		return err
	}
	RzpWalletID = RZPWallet.ID
	amount := RazorpayBalance - RZPWallet.Balance
	wallet.UpdateWalletBalance(RzpWalletID, int64(amount))
	return nil
}
