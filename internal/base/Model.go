package base

type Model struct {
	ID        string `gorm:"type:varchar(20)"`
	CreatedAt int64
	UpdatedAt int64
	DeletedAt int64
}
