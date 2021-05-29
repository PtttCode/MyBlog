package db

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(20); not null; unique"`
	Password string `gorm:"type:varchar(20); not null"`
	Email    string `gorm:"type:varchar(50); not null; unique"`
	UID      string `gorm:"primaryKey;type:char(36);autoIncrement:false"`
}

type Document struct {
	gorm.Model
	Title   string `gorm:"type:varchar(20)"`
	Article string `gorm:"type:text; not null"`
	DocUid  string `gorm:"type:char(36); not null; unique"`
	Privacy int8   `gorm:"type:int8;default:0"`
	UserUID string `gorm:"primaryKey;type:char(36);autoIncrement:false"`
}
