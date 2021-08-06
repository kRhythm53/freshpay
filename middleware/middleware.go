package middleware

import (
	"errors"
	"github.com/freshpay/internal/entities/user_management/session"
	"github.com/gin-gonic/gin"
	"time"
)
func Authenticate(c *gin.Context){
	if c.FullPath() =="/users/signin" || c.FullPath()=="/users/signup" {
		c.Next()
		return
	}
	sessionId:= c.Request.Header["Session_id"][0]

	//
	var Session session.Detail
	err1:=session.GetSessionById(&Session, sessionId)
	if err1!=nil{
		c.AbortWithStatus(403)
		return
	} else if Session.ExpireTime < uint64(time.Now().Unix()){
		c.AbortWithError(400,errors.New("Session has expired"))
		return
	}
	userId:=Session.UserId
	c.Set("userId",userId)


	//
	c.Next()
}
