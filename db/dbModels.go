package db

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct{
	gorm.Model
	//Id	uint	`gorm: "primary_key; AUTO_INCREMENT"`
	Name	string	`gorm: "type:varchar(20); unique; not null"`
	Password	string 	`gorm: "type:varchar(20); not null"`
	Email	string 	`gorm: "type:varchar(50); not null"`
}

type Document struct {
	gorm.Model
	Article	string
	Author string
	User	User	`gorm: "foreginkey: u_id;association_foreginkey:id"`
	UserId	uint	`gorm: "column: uid"`
	Timestamp	time.Time
}