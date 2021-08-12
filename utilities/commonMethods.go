package utilities

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
func CreateID(prefix string,length int) string{
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return prefix+"_"+string(b)
}

func IsNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}


func GetEncryption(password string,hash *string)  error{
	hashByte, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err==nil{
		*hash=string(hashByte)
	}
	return err
}

func MatchPassword(password string,hash string) bool{
	err:=bcrypt.CompareHashAndPassword([]byte(hash),[]byte(password))
	return err==nil
}


/*
ValidatePhoneNumber will validate the Phone Number
 */
func ValidatePhoneNumber(phoneNumber string) (err error){
	if len(phoneNumber)!=10 || phoneNumber[0]=='0'{
		err=errors.New("phone number should be 10 digit long")
		return err
	}
	if !IsNumeric(phoneNumber){
		err=errors.New("Phone number can contain characters 0-9")
		return err
	}
	return nil
}


//CreateOTP
func CreateOTP(otp_length int) string {
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	otp:=strconv.Itoa(seededRand.Intn(int(math.Pow(10,float64(otp_length)))))
	return otp
}


//ValidateBankAccountNumber
func ValidateBankAccountNumber(AccountNumber string) (err error){
	if len(AccountNumber)<9 || len(AccountNumber)>18{
		return errors.New("Number of characters in account number should be b/w 9 and 18")
	}
	return nil
}

//ValidateIFSCCode
func ValidateIFSCCode(IFSCCode string) (err error){
	if len(IFSCCode) !=11{
		return errors.New("Number of characters in IFSCCode should be 11")
	}
	return nil
}