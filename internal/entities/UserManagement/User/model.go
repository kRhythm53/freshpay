package User

type Detail struct {
	EntityId string `json:"entity_id"`
	Name string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Password string `json:"password"`
	Email string `json:"email"`
}

const (
	TableName="user"
	EntityName="user"
)
func(sd *Detail) TableName() string{
	return TableName
}