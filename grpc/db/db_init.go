package db

import (
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"umbrella.github.com/go-micro.example/grpc/entry"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(sqlite.Open(""), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&entry.User{})

	// Create
	db.Create(&entry.User{
		UserId:   100003,
		UserName: "byteew",
		UserPwd:  "zcj123",
		UserDate: time.Now(),
	})

	db.Create(&entry.User{
		UserId:   100004,
		UserName: "kuaishou",
		UserPwd:  "huawei1",
		UserDate: time.Now(),
	})

	db.Create(&entry.User{
		UserId:   100005,
		UserName: "umbrella",
		UserPwd:  "zhangwe",
		UserDate: time.Now(),
	})

	// Write
	var user entry.User
	db.First(&user, 100003)
	fmt.Println(user)
}

func GetDB() *gorm.DB {
	return db
}
