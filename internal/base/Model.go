package base

type Model struct {
	ID        string `gorm:"type:varchar(20)"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	//DeletedAt int64
}
