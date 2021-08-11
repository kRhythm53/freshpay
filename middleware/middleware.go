package middleware

import (
	"github.com/freshpay/internal/entities/Error"
	"github.com/freshpay/internal/entities/admin/admin_session"
	"github.com/freshpay/internal/entities/user_management/user_session"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)


/*
PublicPath stores the public paths, paths that can be accessed by anyone
 */
var PublicPath = []string{
	"/users/signup",
	"/users/signin",
	"/admin/signup",
	"/admin/signin",
	"/users/signup/otp/verification",
	"/admin/signup/otp/verification",
	"/wallet/:phone_number",
}

/*
	isPublicPath will return true if the path is Public else it will return false
 */
func isPublicPath(Path string) bool {
	for _, path := range PublicPath {
		if Path == path {
			return true
		}
	}
	return false
}

/*
userPath will stores the path belonging to user
 */
var userPath = []string{
	"/users/bankaccount",
	"/users/bankaccounts",
	"/users/beneficiary",
	"/users/balance",
	"/payments",
	"/payments/:payments_id",
	"/payments/",
	"/users/complaint",
	"/campaigns/active",
}


/*
adminPath will store path belonging to admin
 */
var adminPath = []string{
	"/admin/complaint/:complaint_id",
	"/admin/complaints",
	"/admin/active_complaints",
	"/campaigns/",
	"/campaigns/:campaign_id",
}

/*
	return if a path belongs to user or not
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
return if a path belongs to admin or not
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
	if isPublicPath(c.FullPath()) {
		c.Next()
		return
	}

	sessionId:=ExtractSessionIdFromHeaders(c)

	if sessionId==""{
		c.JSON(401, Error.Detail{
			"BAD_REQUEST_ERROR","Failed","Invalid Session Id,Please Sign in Again",
			"business","No Session Id in headers","NA","{}",
		})
		c.Abort()
		return
	}

	//if sessionId belongs to user
	sender := strings.Split(sessionId, "_")[0]
	err:=ValidatePath(sender,c.FullPath())
	if err!=nil{
		c.JSON(401,&err)
		c.Abort()
		return
	}

	if sender == user_session.Prefix {
		var Session user_session.Detail
		err=ValidateUserSessionId(sessionId,&Session)
		if err!=nil {
			c.JSON(401,&err)
			c.Abort()
			return
		}

		//set the user Id to gin.Context
		userId := Session.UserId
		c.Set("userId", userId)
		c.Next()
	} else if sender == admin_session.Prefix {
		var Session admin_session.Detail
		err=ValidateAdminSessionId(sessionId,&Session)
		if err!=nil {
			c.JSON(401,&err)
			c.Abort()
			return
		}

		//set admin Id to gin.Context
		adminId := Session.AdminId
		c.Set("adminId", adminId)
		c.Next()
	}

}


/*
ExtractSessionIdFromHeaders will get session Id from the headers
 */
func ExtractSessionIdFromHeaders(c *gin.Context) string{
	header:= c.Request.Header["Session_id"] //S of Session_id has to be capital
	if len(header)==0{
		return ""
	}
	return header[0]
}

/*
ValidatePath  will validate the path of the sender or will error a Error
 */
func ValidatePath(sender string,Path string)  *Error.Detail{
	if (sender==user_session.Prefix && isUserPath(Path)) ||(sender==admin_session.Prefix && isAdminPath(Path)){
		return nil
	}
	err:=Error.Detail{
		"Unauthorized","Failed","Path is not Authorized",
		"business","Access Denied","NA","{}",
	}
	return &err
}


/*
 	ValidateUserSessionId will validate the user session or will return the Rrror
 */
func ValidateUserSessionId(sessionId string,Session *user_session.Detail) *Error.Detail{
	var err error
	err=user_session.GetSessionById(Session,sessionId)
	if err!=nil{
		return &Error.Detail{
			"UnAuthorized","Failed",err.Error(),
			"business","Session Id is Invalid","NA","{}",
		}
	}

	//checking that session hasn't expired
	if Session.ExpireTime < uint64(time.Now().Unix()){
		return &Error.Detail{
			"UnAuthorized","Failed","Session has expired, Please Sign in Again",
			"business","Session has expired","NA","{}",
		}
	}
	return nil
}

/*
	ValidateAdminSessionId will validate the user session or will return the Rrror
*/
func ValidateAdminSessionId(sessionId string,Session *admin_session.Detail) *Error.Detail{
	var err error
	err=admin_session.GetSessionById(Session,sessionId)
	if err!=nil{
		return &Error.Detail{
			"UnAuthorized","Failed",err.Error(),
			"business","Session Id is Invalid","NA","{}",
		}
	}

	//checking that session hasn't expired
	if Session.ExpireTime < uint64(time.Now().Unix()){
		return &Error.Detail{
			"UnAuthorized","Failed","Session has expired, Please Sign in Again",
			"business","Session has expired","NA","{}",
		}
	}
	return nil
}