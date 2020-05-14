package db

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct{
	gorm.Model
	Username string `gorm:"type:varchar(20); unique; not null"`
	Password string `gorm:"type:varchar(20); not null"`
	Email    string `gorm:"type:varchar(50); not null; unique"`
}

type Document struct {
	gorm.Model
	Title	string	`gorm:"type:varchar(20)"`
	Article	string	`gorm:"type:text; not null"`
	Author string	`gorm:"type:varchar(30); unique"`
	Uid	uint
	Timestamp	time.Time
}