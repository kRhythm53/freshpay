package middleware

import (
	"errors"
	"fmt"
	"github.com/freshpay/internal/entities/admin/admin_session"
	"github.com/freshpay/internal/entities/user_management/session"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)
var userPath=[]string{
	"/users/bankaccount",
	"/users/bankaccounts",
	"/users/beneficiary",
	"/users/complaint",
}

var adminPath=[]string{
	"/admin/complaint/:complaint_id",
	"/admin/complaints",
	"/admin/active_complaints",
}


/*
	return if a method belongs to user or not
*/
func isUserPath(Path string) bool{
	for _,path:=range userPath{
		if Path==path{
			return true
		}
	}
	return false
}

/*
return if a method belongs to admin or not
 */
func isAdminPath(Path string) bool{
	for _,path:=range adminPath{
		if Path==path{
			return true
		}
	}
	return false
}



func Authenticate(c *gin.Context){
	fmt.Println("c.Path: ",c.FullPath())
	if c.FullPath()=="/users/signup" || c.FullPath()=="/admin/signup"{
		fmt.Println("Inside")
		/*err:= phonenumber_verification.VerifyPhoneNumber(c)
		if err!=nil{
			c.AbortWithError(400,err)
			return
		} else{
			c.Next()
			return
		}*/
		c.Next()
		return
	}
	if c.FullPath() =="/users/signin"  || c.FullPath() =="/admin/signin"  {
			c.Next()
			return
	}

	sessionId:= c.Request.Header["Session_id"][0]
	//if sessionId belongs to user
	sender:=strings.Split(sessionId,"_")[0]
	if sender==session.Prefix{
		if !isUserPath(c.FullPath()){
			c.AbortWithError(400, errors.New("acess denied"))
			return
		}
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
	} else if sender==admin_session.Prefix{
		println(c.FullPath())
		if !isAdminPath(c.FullPath()){
			c.AbortWithError(400, errors.New("acess denied"))
			return
		}
		var Session admin_session.Detail
		err1:=admin_session.GetSessionById(&Session, sessionId)
		if err1!=nil{
			c.AbortWithStatus(403)
			return
		} else if Session.ExpireTime < uint64(time.Now().Unix()){
			c.AbortWithError(400,errors.New("Session has expired"))
			return
		}
		adminId:=Session.AdminId
		c.Set("adminId",adminId)
	}
	c.Next()
}
