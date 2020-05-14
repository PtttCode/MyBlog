package db

import (
	st "../settings"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"strings"
	"time"
)

var db *gorm.DB

func GetDB() *gorm.DB{
	return db
}

func DBInit(){
	dbPath := strings.Join([]string{st.MysqlUser, ":", st.MysqlPasswd, "@(", st.MysqlIp, ":", st.MysqlPort, ")/", st.MysqlDatabase, "?charset=utf8&parseTime=true"}, "")
	db_tmp, err := gorm.Open("mysql", dbPath)
	if err != nil{
		panic(err)
	}
	db = db_tmp
	db.SingularTable(true)
	db.DB().SetConnMaxLifetime(1 * time.Second)
	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(2000)
	db.SingularTable(true)
	db.LogMode(true)

	//db.AutoMigrate(
	//	&User{Username: "pt", Password: "123456", Email:"tongpang@webot.co"},
	//	&Document{},
	//)
	//db.Create(&User{Username: "pt", Password: "123456", Email:"tongpang@webot.co"})
	//db.Create(&Document{})
	//db.DropTable(&User{})
	//db.DropTable(&Document{})
	db.CreateTable(&User{})
	db.CreateTable(&Document{})
	//db.Model(&Document{}).AddForeignKey("uid", "user(id)", "RESTRICT", "RESTRICT")
}

func GetUserId(username string) uint{
	var user User
	q := db.Where("username = ?", username).First(&user)
	if q.Error != nil{
		return 0
	}
	return user.ID
}
