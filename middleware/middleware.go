package middleware

import (
	"errors"
	"fmt"
	"github.com/freshpay/internal/entities/admin/admin_session"
	"github.com/freshpay/internal/entities/user_management/user_session"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

var noSessionIdPath = []string{
	"/users/signup",
	"/users/signin",
	"/admin/signup",
	"/admin/signin",
	"/users/signup/otp/verification",
	"/admin/signup/otp/verification",
	"/wallet/:phone_number",
}

func isNoSessionIdPath(Path string) bool {
	fmt.Println(Path)
	for _, path := range noSessionIdPath {
		if Path == path {
			return true
		}
	}
	return false
}

var userPath = []string{
	"/users/bankaccount",
	"/users/bankaccounts",
	"/users/beneficiary",
	"/payments",
	"/payments/:payments_id",
	"/payments/",
	"/campaigns/",
	"/campaigns/:campaign_id",
	"/users/complaint",
}

var adminPath = []string{
	"/admin/complaint/:complaint_id",
	"/admin/complaints",
	"/admin/active_complaints",
}

/*
	return if a method belongs to user or not
*/
func isUserPath(Path string) bool {
	for _, path := range userPath {
		if Path == path {
			return true
		}
	}
	return false
}

/*
return if a method belongs to admin or not
*/
func isAdminPath(Path string) bool {
	for _, path := range adminPath {
		if Path == path {
			return true
		}
	}
	return false
}

func Authenticate(c *gin.Context) {
	if isNoSessionIdPath(c.FullPath()) {
		c.Next()
		return
	}
	if len(c.Request.Header["Session_id"])==0{
		c.AbortWithError(403,errors.New("Session Id is invalid"))
		return
	}
	sessionId := c.Request.Header["Session_id"][0]
	if len(sessionId) <user_session.IDLengthExcludingPrefix{
		c.AbortWithError(403,errors.New("Session Id is invalid"))
		return
	}
	//if sessionId belongs to user
	sender := strings.Split(sessionId, "_")[0]
	if sender == user_session.Prefix {
		if !isUserPath(c.FullPath()) {
			c.AbortWithError(403, errors.New("acess denied"))
			return
		}
		var Session user_session.Detail
		err1 := user_session.GetSessionById(&Session, sessionId)
		if err1 != nil {
			c.AbortWithError(403,errors.New("Session Id is invalid, Please Sign in again"))
			return
		} else if Session.ExpireTime < uint64(time.Now().Unix()) {
			c.AbortWithError(400, errors.New("Session has expired , Please Sign in again"))
			return
		}
		userId := Session.UserId
		c.Set("userId", userId)
	} else if sender == admin_session.Prefix {
		if !isAdminPath(c.FullPath()) {
			c.AbortWithError(403, errors.New("acess denied"))
			return
		}
		var Session admin_session.Detail
		err1 := admin_session.GetSessionById(&Session, sessionId)
		if err1 != nil {
			c.AbortWithError(403,errors.New("Session Id is invalid, Please Sign in again"))
			return
		} else if Session.ExpireTime < uint64(time.Now().Unix()) {
			c.AbortWithError(400, errors.New("Session has expired, Please Sign in again"))
			return
		}
		adminId := Session.AdminId
		c.Set("adminId", adminId)
	}
	c.Next()
}
