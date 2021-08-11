package campaigns
import (
	"fmt"
	"github.com/freshpay/internal/entities/campaigns"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCampaign(c *gin.Context) {
	var user [] campaigns.Campaign
	err := campaigns.GetAllCampaigns(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func CreateCampaign(c *gin.Context) {
	var user campaigns.Campaign
	err := c.BindJSON(&user)
	if err != nil {
		fmt.Println("could not read json")
		return
	}
	fmt.Println(&user)
	err2 := campaigns.CreateCampaign(&user)
	if err2 != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func GetCampaignByID(c *gin.Context) {
	id := c.Params.ByName("campaign_id")
	var user campaigns.Campaign
	err := campaigns.GetCampaignByID(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func UpdateCampaign(c *gin.Context) {
	var user campaigns.Campaign
	id := c.Params.ByName("campaign_id")
	err := campaigns.GetCampaignByID(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound,user)
	}
	err = c.BindJSON(&user)
	if err !=nil {
		return
	}
	err = campaigns.UpdateCampaign(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
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