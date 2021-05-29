package db

import (
	"strings"
	"time"

	st "MyBlog/settings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var dbHandle *gorm.DB

func GetDB() *gorm.DB {
	return dbHandle
}

func DBInit() {
	dbPath := strings.Join([]string{st.MysqlUser, ":", st.MysqlPasswd, "@(", st.MysqlIp, ":", st.MysqlPort, ")/", st.MysqlDatabase, "?charset=utf8&parseTime=true"}, "")
	db_tmp, err := gorm.Open("mysql", dbPath)
	if err != nil {
		panic(err)
	}
	dbHandle = db_tmp
	dbHandle.SingularTable(true)
	dbHandle.DB().SetConnMaxLifetime(1 * time.Second)
	dbHandle.DB().SetMaxIdleConns(20)
	dbHandle.DB().SetMaxOpenConns(2000)
	dbHandle.LogMode(true)

	dbHandle.AutoMigrate(
		&User{},
		&Document{},
	)
	//db.Create(&User{Username: "pt", Password: "123456", Email:"tongpang@webot.co"})
	//db.Create(&Document{})
	//db.DropTable(&User{})
	//db.DropTable(&Document{})
}

func GetUserId(username string) uint {
	var user User
	q := dbHandle.Where("username = ?", username).First(&user)
	if q.Error != nil {
		return 0
	}
	return user.ID
}
