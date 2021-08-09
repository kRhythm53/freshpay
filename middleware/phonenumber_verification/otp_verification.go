package phonenumber_verification

import (
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/souvikhaldar/gobudgetsms"
	"io"
	"time"
)

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
var otp_length = 6
var BsmsUsername = "Rhythm"
var BsmsUserId = "22210"
var BsmsHandle = "0001e6b45069af5ebcd5d1fc7a975e86"

func EncodeToString(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}


var redisClient = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

// SetValue sets the key value pair
func SetValue(key string, value string, expiry time.Duration) error {
	err := redisClient.Set(key, value, expiry).Err()
	if err != nil {
		fmt.Println("can't set the value")
		return err
	}
	return nil
}


var smsConfig gobudgetsms.Details
func init() {
	smsConfig = gobudgetsms.SetConfig(BsmsUsername, BsmsUserId, BsmsHandle, "", 1, 0, 0)

}
// sendmessage will send sms using gobudgetsms
func sendmessage(phoneNumber string, otp string) error{
	message := "Your OTP for freshpay signup is " + otp
	res, err := gobudgetsms.SendSMS(smsConfig, message,phoneNumber , "freshpay")
	if err != nil {
		return err
	}
	fmt.Println("The response after sending sms is ", res)
	return nil
}



func Verify(phoneNumber string)(err error){
	otp := EncodeToString(otp_length)
	if err= SetValue(phoneNumber, otp, 5*time.Minute); err != nil {
		err=errors.New("Error in setting value to redis")
		return err
	}
	if err := sendmessage(phoneNumber, otp); err!= nil{
		err=errors.New("Could not sent otp message")
		return err
	}
	return err
}