package model

type User struct {
	Id        int64  `json:"id"`
	Username  string `json:"user_name"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	IsDeleted int    `json:"is_deleted"`
}

func (User) TableName() string {
	return "user"
}
