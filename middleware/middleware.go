package middleware

import (
	"github.com/freshpay/internal/entities/admin/admin_session"
	"github.com/freshpay/internal/entities/user_management/user_session"
	"github.com/gin-gonic/gin"
	"net/http"
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
	"/users/balance",

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
		c.JSON(401, gin.H{
			"Code": "BAD_REQUEST_ERROR",
			"Status":"Failed",
			"Description":"Invalid Session Id,Please Sign in Again",
			"Source": "business",
			"Reason": "No Session Id in headers",
			"Step": "NA",
			"Metadata":"{}",
		})
		c.Abort()
		return
	}
	sessionId := c.Request.Header["Session_id"][0]
	if len(sessionId) <user_session.IDLengthExcludingPrefix{
		c.JSON(401, gin.H{
			"Code": "BAD_REQUEST_ERROR",
			"Status":"Failed",
			"Description":"Invalid Session Id,Please Sign in Again",
			"Source": "business",
			"Reason": "Invalid Session ID",
			"Step": "NA",
			"Metadata":"{}",
		})
		c.Abort()
		return
	}
	//if sessionId belongs to user
	sender := strings.Split(sessionId, "_")[0]
	if sender == user_session.Prefix {
		if !isUserPath(c.FullPath()) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"Code": "Unauthorized",
				"Status":"Failed",
				"Description":"Access Denied",
				"Source": "business",
				"Reason": "Path is not Authorized",
				"Step": "NA",
				"Metadata":"{}",
			})
			c.Abort()
		}else{
			var Session user_session.Detail
			err1 := user_session.GetSessionById(&Session, sessionId)
			if err1 != nil {
				c.JSON(403, gin.H{
					"Code": "Unauthorized",
					"Status":"Failed",
					"Description":"Session Id is invalid, Please Signin Again",
					"Source": "business",
					"Reason": "Session Id is invalid",
					"Step": "NA",
					"Metadata":"{}",
				})
				c.Abort()
				return
			} else if Session.ExpireTime < uint64(time.Now().Unix()) {
				c.JSON(403, gin.H{
					"Code": "Unauthorized",
					"Status":"Failed",
					"Description":"Session has expired, Please Signin Again",
					"Source": "business",
					"Reason": "Session has expired",
					"Step": "NA",
					"Metadata":"{}",
				})
				c.Abort()
				return
			}else{
				userId := Session.UserId
				c.Set("userId", userId)
				c.Next()
			}
		}
	} else if sender == admin_session.Prefix {
		if !isAdminPath(c.FullPath()) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"Code": "Unauthorized",
				"Status":"Failed",
				"Description":"Access Denied",
				"Source": "business",
				"Reason": "Path is not Authorized",
				"Step": "NA",
				"Metadata":"{}",
			})
			c.Abort()
		} else{
			var Session admin_session.Detail
			err1 := admin_session.GetSessionById(&Session, sessionId)
			if err1 != nil {
				c.JSON(403, gin.H{
					"Code": "Unauthorized",
					"Status":"Failed",
					"Description":"Session Id is invalid, Please Signin Again",
					"Source": "business",
					"Reason": "Session Id is invalid",
					"Step": "NA",
					"Metadata":"{}",
				})
				c.Abort()
				return
			} else if Session.ExpireTime < uint64(time.Now().Unix()) {
				c.JSON(403, gin.H{
					"Code": "Unauthorized",
					"Status":"Failed",
					"Description":"Session has expired, Please Signin Again",
					"Source": "business",
					"Reason": "Session has expired",
					"Step": "NA",
					"Metadata":"{}",
				})
				c.Abort()
				return
			} else{
				adminId := Session.AdminId
				c.Set("adminId", adminId)
				c.Next()
			}
		}

	}

}
