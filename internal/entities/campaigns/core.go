package campaigns

import (
	"github.com/freshpay/internal/config"
	"github.com/freshpay/internal/entities/payments/utilities"
	"github.com/freshpay/internal/entities/user_management/user"
	"math"
)

func GetAllCampaigns(campaign *[]Campaign ) (err error) {
	if err = config.DB.Find(campaign).Error; err != nil {
		return err
	}
	return nil
}

func CreateCampaign(campaign *Campaign  ) (err error) {
	campaign.ID =  utilities.RandomString(14,ComplaintPrefix)
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
	config.DB.Save(campaign)
	return nil
}

func DeleteCampaign(campaign *Campaign, id string) (err error) {
	config.DB.Table("campaign").Where("id = ?", id).Delete(campaign)
	return nil
}
func Eligibility (Time int64,Amount int64,userid string) int  {
	var UserRow user.Detail
	cashback:=0
	err1:=config.DB.Table("user").Where("id = ?",userid).First(&UserRow).Error
	if err1 != nil {
		return cashback
	}
	TransNum:= UserRow.NumberOfTransactions
	var users [] Campaign
	//Time:=payment.CreatedAt
	//Amount:=payment.Amount
	err := config.DB.Table("campaign").Where("start_time <= ? AND end_time >= ? AND transaction_number = ?",Time,Time,TransNum).Find(&users).Error
	if err != nil {
		return cashback
	} else {
		for _,entry:=range users {
			if entry.IsActive {
				percentage:=entry.PercentageRate
				PercentageAmount:= (float64(percentage))*(float64(Amount))/100
				cashbackAmount:= math.Min(PercentageAmount, float64(entry.MaxCashback))
				cashback= int(math.Max(float64(cashback), cashbackAmount))
			}
		}
	}
	return cashback
}