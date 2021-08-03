package Bank
import "github.com/kshitij-nawandar9/freshpay/internal/Config"

//CreateBank will create a new bank
func CreateBank(bank *Detail)(err error){
	if err=Config.DB.Create(bank).Error; err!=nil{
		return err
	}
	return nil
}

//GetBankById will return the bank by using the id
func GetBankByID(bank *Detail, entity_id string)(err error){
	if err = Config.DB.Where("entity_id = ?", entity_id).First(bank).Error; err != nil {
		return err
	}
	return nil
}


