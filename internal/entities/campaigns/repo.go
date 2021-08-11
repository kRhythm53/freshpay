package campaigns

import (
	"github.com/freshpay/internal/config"
	"github.com/freshpay/internal/entities/user_management/user"
)

func CreateCampaignToDB(campaign *Campaign) (err error) {
	if err := config.DB.Table("campaign").Create(campaign).Error; err != nil {
		return err
	}
	return nil
}

func GetCampaignByIDFromDB(campaign *Campaign, id string) (err error) {
	if err = config.DB.Table("campaign").Where("id = ?", id).First(campaign).Error; err != nil {
		return err
	}
	return nil
}

func GetAllCampaignsFromDB(campaign *[]Campaign) (err error) {
	if err = config.DB.Table("campaign").Where("is_active = ?", true).Find(campaign).Error; err != nil {
		return err
	}
	return nil
}
func UpdateCampaignFromDB(campaign *Campaign) (err error) {
	config.DB.Table("campaign").Save(campaign)
	return nil
}
func DeleteCampaignFromDB(campaign *Campaign, id string) (err error) {
	config.DB.Table("campaign").Where("id = ?", id).Delete(campaign)
	return nil
}
func UserQueryFromDB(UserRow *user.Detail, userid string) (err error) {
	return config.DB.Table("user").Where("id = ?", userid).First(&UserRow).Error
}
func ValidCampaignQuery(Time int64, TransNum int64, ValidCampaigns *[]Campaign) (err error) {
	return config.DB.Table("campaign").Where("start_time <= ? AND end_time >= ? AND transaction_number = ?", Time, Time, TransNum).Find(ValidCampaigns).Error
}
