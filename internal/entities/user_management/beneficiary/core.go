package beneficiary

import (
	"github.com/freshpay/internal/config"
	"github.com/freshpay/internal/entities/user_management/utilities"
)

//CreateBeneficiary will create a new beneficiary
func CreateBeneficiary(beneficiary *Detail,userId string)(err error){
	beneficiary.ID=utilities.CreateID(14)
	beneficiary.UserId=userId
	if err=config.DB.Create(beneficiary).Error; err!=nil{
		return err
	}
	return nil
}

//GetBeneficiaryById will return the beneficiary by using the id
func GetBeneficiaryById(beneficiary *Detail,id string)(err error){
	if err = config.DB.Where("ID = ?", id).First(beneficiary).Error; err != nil {
		return err
	}
	return nil
}


//GetAllBeneficiaryByUserId will return all the beneficiary attached to a user
func GetAllBeneficiaryAccountsByUserId(beneficiary *[]Detail,user_id string)(err error){
	if err=config.DB.Where("user_id = ?", user_id).Find(beneficiary).Error; err!=nil{
		return err
	}
	return nil
}