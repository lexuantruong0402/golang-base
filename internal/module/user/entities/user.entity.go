package entities

type User struct {
	Id        int    `json:"id" gorm:"primaryKey"`
	Username  string `json:"username" gorm:"unique"`
	Password  string `json:"password"`
	Email     string `json:"email" gorm:"unique"`
	CreatedAt string `json:"created_at" gorm:"autoCreateTime:true"`
	UpdatedAt string `json:"updated_at" gorm:"autoUpdateTime:false"`
}
