package Session

type Detail struct {
	EntityId string `json:"entity_id"`
	UserId string `json:"user_id"`
	ExpireTime uint64 `json:"expire_time"`
}

const (
	TableName="session"
	EntityName="session"
)

func(sd *Detail) TableName() string{
	return TableName
}