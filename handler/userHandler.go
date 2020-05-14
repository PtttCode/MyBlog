package handler

import (
	"../db"
	"fmt"
	"github.com/kataras/golog"
)

type LoginForm struct{
	Username string	`form:"username"`
	PassWord string	`form:"password"`
}

type RegisterForm struct {
	Username string `form:"username"`
	Email string `form:"email"`
	Password string `form:"password"`
	MobileNum string `form:"phoneNum"`
}

type EasyResponseJson struct {
	Code int `json:"code"`
	Massage string `json:"massage"`
}

func UserRegister(form RegisterForm) EasyResponseJson{
	fmt.Println(form)
	var middleUser db.User
	if form.Username != "" && form.Password != "" && form.Email != ""{
		dbHandler := db.GetDB()
		var user = db.User{Username: form.Username, Password: form.Password, Email: form.Email}

		query := dbHandler.Where("name = ?", user.Username).Or("email = ?", user.Email).First(&middleUser)
		if query.RowsAffected != 0{
			return EasyResponseJson{-1, "用户重复，请重新填写"}
		}
		res := dbHandler.Create(&user)
		if res.Error != nil {
			golog.Error(res.Error)
			return EasyResponseJson{-1, "创建失败"}
		}
	}else{
		return EasyResponseJson{-1, "请填写必要信息"}
	}
	return EasyResponseJson{0, "注册成功"}
}

func UserLogin(form LoginForm) string {
	fmt.Println(form)
	return "Hello " + form.Username
}

