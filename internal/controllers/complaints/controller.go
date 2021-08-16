package complaints

import (
	"github.com/freshpay/internal/base"
	"github.com/freshpay/internal/entities/complaints"
	"github.com/gin-gonic/gin"
	"net/http"
)


type Response struct{
	Entity string
	Complaint complaints.Complaint
}
func CreateComplaint(c *gin.Context) {
	var Complaint complaints.Complaint
	err := c.BindJSON(&Complaint)
	if err != nil {
		c.JSON(http.StatusBadRequest,base.Failure{Error: base.Error{Code: "Bad request error", Description: err.Error(), Source: "business", Reason: "validation failed", Step: "NA"}})
		return
	}
	userId :=c.GetString("userId")
	err = complaints.CreateComplaint(&Complaint,userId)
	if err != nil {
		c.JSON(http.StatusBadRequest,base.Failure{Error: base.Error{Code: "Bad request error", Description: err.Error(), Source: "business", Reason: "validation failed", Step: "NA"}})
	} else {
		resp:= Response{Entity: "Complaints",Complaint: Complaint}
		c.JSON(http.StatusOK, resp)
	}
}


func UpdateComplaintById(c *gin.Context) {
	var Complaint complaints.Complaint
	id := c.Params.ByName("complaint_id")
	adminId := c.GetString("adminId")
	var Refund complaints.Refund
	c.BindJSON(&Refund)
	refund := Refund.Refund
	err := complaints.GetComplaintByID(&Complaint, id)
	if err != nil {
		c.JSON(http.StatusBadRequest,base.Failure{Error: base.Error{Code: "Bad request error", Description: err.Error(), Source: "business", Reason: "validation failed", Step: "NA"}})
	} else{
		c.BindJSON(&Complaint)
		err = complaints.UpdateComplaint(&Complaint, id, refund,adminId)
		if err != nil {
			c.JSON(http.StatusBadRequest,base.Failure{Error: base.Error{Code: "Bad request error", Description: err.Error(), Source: "business", Reason: "validation failed", Step: "NA"}})
		} else {
			resp:= Response{Entity: "Complaints",Complaint: Complaint}
			c.JSON(http.StatusOK, resp)
		}
	}
}


func GetComplaints(c *gin.Context){
	var Complaint []complaints.Complaint
	err := complaints.GetAllComplaints(&Complaint)
	if err != nil {
		c.JSON(http.StatusBadRequest,base.Failure{Error: base.Error{Code: "Bad request error", Description: err.Error(), Source: "business", Reason: "validation failed", Step: "NA"}})
	} else {
		resp := make([]Response, len(Complaint))
		for i,payment := range Complaint{
			resp[i]= Response{Entity: "Complaints",Complaint: payment}
		}
		c.JSON(http.StatusOK, resp)
	}
}

func GetActiveComplaints(c *gin.Context){
	var Complaint []complaints.Complaint
	err := complaints.GetAllActiveComplaints(&Complaint)
	if err != nil {
		c.JSON(http.StatusBadRequest,base.Failure{Error: base.Error{Code: "Bad request error", Description: err.Error(), Source: "business", Reason: "validation failed", Step: "NA"}})
	} else {
		resp := make([]Response, len(Complaint))
		for i,payment := range Complaint{
			resp[i]= Response{Entity: "Complaints",Complaint: payment}
		}
		c.JSON(http.StatusOK, resp)
	}
}

func GetComplaintById(c *gin.Context) {
	id := c.Params.ByName("complaint_id")
	var Complaint complaints.Complaint
	err := complaints.GetComplaintByID(&Complaint, id)
	if err != nil {
		c.JSON(http.StatusBadRequest,base.Failure{Error: base.Error{Code: "Bad request error", Description: err.Error(), Source: "business", Reason: "validation failed", Step: "NA"}})
	} else {
		resp:= Response{Entity: "Complaints",Complaint: Complaint}
		c.JSON(http.StatusOK, resp)
	}
}