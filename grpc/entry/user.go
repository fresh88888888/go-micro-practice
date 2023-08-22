package entry

import "time"

type User struct {
	UserId   int       `gorm:"column:user_id;AUTO_INCREMEMT;PRIMARY_KEY"`
	UserName string    `gorm:"column:user_name;type;varchar(50);unique_index"`
	UserPwd  string    `gorm:"column:user_pwd;type:varchar(20)"`
	UserDate time.Time `gorm:"column:user_date"`
}
