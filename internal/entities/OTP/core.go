package OTP

import (
	"errors"
	"fmt"
	"github.com/freshpay/utilities"
	"github.com/souvikhaldar/gobudgetsms"
	"time"
)


/*
will create a random number
*/



// SetValue sets the key value pair
func SetValue(key string, value string, expiry time.Duration) error {
	err := redisClient.Set(key, value, expiry).Err()
	return err
}



// sendmessage will send sms using gobudgetsms
func sendmessage(phoneNumber string, otp string) error{
	return nil //need to remove this line to send message
	message := "Your OTP for freshpay signup is " + otp
	res, err := gobudgetsms.SendSMS(smsConfig, message,phoneNumber , "freshpay")
	if err != nil {
		return err
	}
	fmt.Println("The response after sending sms is ", res)
	return nil
}

/*
	this function will create otp and will send the otp and save the otp
*/
func SendOTP(phoneNumber string)(err error){
	otp:=utilities.CreateOTP(otp_length)
	err=SetValue(phoneNumber,otp,ExpireTime)
	if err!=nil{
		//err=errors.New("Error in setting OTP to redis")
		return err
	} else{
		err=sendmessage(phoneNumber,otp)
		if err!=nil{
			return err;
		}
	}
	return nil
}

func VerifyOTP(otp Detail)(err error){
	OTP,err:=redisClient.Get(otp.PhoneNumber).Result()
	if err!=nil{
		return err
	}
	if OTP!=otp.OTP{
		return errors.New("OTP did not Match")
	}
	return err
}