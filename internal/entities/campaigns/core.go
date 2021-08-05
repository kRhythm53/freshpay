package campaigns

import (
	"fmt"
	"github.com/freshpay/internal/config"
	"github.com/freshpay/internal/entities/payments/utilities"
)

func GetAllCampaigns(campaign *[]Campaign ) (err error) {
	if err = config.DB.Find(campaign).Error; err != nil {
		return err
	}
	return nil
}

func CreateCampaign(campaign *Campaign  ) (err error) {
	campaign.ID ="cmplt_"+utilities.RandomString(14)
	if err = config.DB.Table("campaign").Create(campaign).Error; err != nil {
		return err
	}
	return nil
}

func GetCampaignByID(campaign *Campaign , id string) (err error) {
	if err = config.DB.Table("campaign").Where("id = ?", id).First(campaign).Error; err != nil {
		return err
	}
	return nil
}

func UpdateCampaign(campaign *Campaign ) (err error) {
	fmt.Println(campaign)
	config.DB.Save(campaign)
	return nil
}

func DeleteCampaign(campaign *Campaign, id string) (err error) {
	config.DB.Table("campaign").Where("id = ?", id).Delete(campaign)
	return nil
}
