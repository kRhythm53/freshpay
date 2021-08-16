package utilities

import (
	"golang.org/x/crypto/bcrypt"
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
