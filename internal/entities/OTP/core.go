package OTP

import (
	"errors"
	"math"
	"math/rand"
	"strconv"
	"time"
)


/*
will create a random number
 */
func CreateOTP() string {
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	otp:=strconv.Itoa(seededRand.Intn(int(math.Pow(10,otp_length))))
	return otp
}


// SetValue sets the key value pair
func SetValue(key string, value string, expiry time.Duration) error {
	err := redisClient.Set(key, value, expiry).Err()
	return err
}



// sendmessage will send sms using gobudgetsms
func sendmessage(phoneNumber string, otp string) error{
	return nil //need to remove this line to send message
}

/*
	this function will create otp and will send the otp and save the otp
 */
func SendOTP(phoneNumber string)(err error){
	otp:=CreateOTP()
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
