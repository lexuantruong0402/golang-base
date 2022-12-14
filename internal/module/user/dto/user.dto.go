package dto

type UserDto struct {
	Id        int    `json:"id" gorm:"primaryKey"`
	Username  string `json:"username" gorm:"unique"`
	Password  string `json:"password"`
	Email     string `json:"email" gorm:"unique"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
