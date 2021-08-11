package campaigns

import (
	"fmt"
	"github.com/freshpay/internal/entities/user_management/user"
	utilities2 "github.com/freshpay/utilities"
	"math"
)

func GetAllCampaigns(campaign *[]Campaign) (err error) {
	return GetAllCampaignsFromDB(campaign)
}

func CreateCampaign(campaign *Campaign) (err error) {
	campaign.ID = utilities2.RandomString(14, CampaignPrefix)
	return CreateCampaignToDB(campaign)
}

func GetCampaignByID(campaign *Campaign, id string) (err error) {
	return GetCampaignByIDFromDB(campaign, id)
}

func UpdateCampaign(campaign *Campaign) (err error) {
	return UpdateCampaignFromDB(campaign)
}

func DeleteCampaign(campaign *Campaign, id string) (err error) {
	return DeleteCampaignFromDB(campaign, id)
}

func Eligibility(Time int64, Amount int64, userid string) int {
	var UserRow user.Detail
	cashback := 0
	err1 := UserQueryFromDB(&UserRow, userid)
	if err1 != nil {
		return cashback
	}
	TransNum := UserRow.NumberOfTransactions
	var ValidCampaigns []Campaign
	err := ValidCampaignQuery(Time, TransNum, &ValidCampaigns)
	if err != nil {
		return cashback
	} else {
		index := -1
		for i, entry := range ValidCampaigns {
			fmt.Println(entry.IsActive, entry.Count)
			if entry.IsActive && entry.Count > 0 {
				percentage := entry.PercentageRate
				PercentageAmount := (float64(percentage)) * (float64(Amount)) / 100
				cashbackAmount := math.Min(PercentageAmount, float64(entry.MaxCashback))
				if cashbackAmount > float64(cashback) {
					cashback = int(math.Max(float64(cashback), cashbackAmount))
					index = i
				}
			}
		}
		if index != -1 {
			issued := &ValidCampaigns[index]
			issued.Count -= 1
			err = UpdateCampaignFromDB(issued)
		}
	}
	return cashback
}