package campaigns

import (
	"github.com/freshpay/internal/base"
	"github.com/freshpay/internal/entities/campaigns"
	"github.com/gin-gonic/gin"
	"net/http"
)
type Response struct{
	Entity string
	Campaign campaigns.Campaign
}
func GetCampaign(c *gin.Context) {
	var campaign [] campaigns.Campaign
	err := campaigns.GetAllCampaigns(&campaign)
	if err != nil {
		c.JSON(http.StatusBadRequest,base.Failure{Error: base.Error{Code: "Bad request error", Description: err.Error(), Source: "business", Reason: "validation failed", Step: "NA"}})
	} else {
		resp := make([]Response, len(campaign))
		for i,payment := range campaign{
			resp[i]= Response{Entity: "Campaigns",Campaign: payment}
		}
		c.JSON(http.StatusOK, resp)
	}
}

func CreateCampaign(c *gin.Context) {
	var user campaigns.Campaign
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest,base.Failure{Error: base.Error{Code: "Bad request error", Description: err.Error(), Source: "business", Reason: "validation failed", Step: "NA"}})
		return
	}
	//fmt.Println(&user)
	err2 := campaigns.CreateCampaign(&user)
	if err2 != nil {
		c.JSON(http.StatusBadRequest,base.Failure{Error: base.Error{Code: "Bad request error", Description: err.Error(), Source: "business", Reason: "validation failed", Step: "NA"}})
	} else {
		resp:=Response{Entity: "Campaigns",Campaign: user}
		c.JSON(http.StatusOK, resp)
	}
}

func GetCampaignByID(c *gin.Context) {
	id := c.Params.ByName("campaign_id")
	var user campaigns.Campaign
	err := campaigns.GetCampaignByID(&user, id)
	if err != nil {
		c.JSON(http.StatusBadRequest,base.Failure{Error: base.Error{Code: "Bad request error", Description: err.Error(), Source: "business", Reason: "validation failed", Step: "NA"}})
	} else {
		resp:=Response{Entity: "Campaigns",Campaign: user}
		c.JSON(http.StatusOK, resp)
	}
}

func UpdateCampaign(c *gin.Context) {
	var user campaigns.Campaign
	id := c.Params.ByName("campaign_id")
	err := campaigns.GetCampaignByID(&user, id)
	if err != nil {
		c.JSON(http.StatusBadRequest,base.Failure{Error: base.Error{Code: "Bad request error", Description: err.Error(), Source: "business", Reason: "validation failed", Step: "NA"}})
	}
	err = c.BindJSON(&user)
	if err !=nil {
		c.JSON(http.StatusBadRequest,base.Failure{Error: base.Error{Code: "Bad request error", Description: err.Error(), Source: "business", Reason: "validation failed", Step: "NA"}})
		return
	}
	err = campaigns.UpdateCampaign(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest,base.Failure{Error: base.Error{Code: "Bad request error", Description: err.Error(), Source: "business", Reason: "validation failed", Step: "NA"}})
	} else {
		resp:=Response{Entity: "Campaigns",Campaign: user}
		c.JSON(http.StatusOK, resp)
	}
}

func DeleteCampaign(c *gin.Context) {
	var user campaigns.Campaign
	id := c.Params.ByName("campaign_id")
	err := campaigns.DeleteCampaign(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}