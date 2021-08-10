package admin_management/*
	CreateProduct will add a new product to the database
*/import (
	"github.com/freshpay/internal/entities/admin"
	"github.com/freshpay/internal/entities/admin/admin_session"
	"github.com/gin-gonic/gin"
	"net/http"
)


func SignUp(c *gin.Context) {
	var Admin admin.Detail
	c.BindJSON(&Admin)
	println(c.BindJSON(&Admin))
	err := admin.SignUp(&Admin)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest,err)
	} else {
		c.JSON(http.StatusOK, Admin)
	}
}


//LoginByPassword will login the admin by Password
func LoginByPassword(c *gin.Context){
	var loginInfo admin.Detail
	c.BindJSON(&loginInfo)
	var Session admin_session.Detail
	err:=admin.LoginByPassword(loginInfo.PhoneNumber,loginInfo.Password,&Session)
	if err!=nil{
		c.AbortWithError(400,err)
	} else{
		c.JSON(http.StatusOK,gin.H{
			"session_id":Session.ID,
		})
	}
}

