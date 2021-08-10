package OTP

import (
	"github.com/go-redis/redis"
	"github.com/souvikhaldar/gobudgetsms"
	"time"
)

type Detail struct {
	PhoneNumber string
	OTP string
}
const (
	otp_length = 6
	BsmsUsername = "Rhythm"
	BsmsUserId = "22210"
	BsmsHandle = "0001e6b45069af5ebcd5d1fc7a975e86"
	ExpireTime = 5*time.Minute


)
var redisClient = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})
var smsConfig = gobudgetsms.SetConfig(BsmsUsername, BsmsUserId, BsmsHandle,
	"", 1, 0, 0)

