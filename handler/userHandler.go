package handler

import (
	"../db"
	"fmt"
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
	if form.Username != "" && form.Password != "" && form.Email != ""{
		dbHandler := db.GetDB()
		var user = db.User{Name: form.Username, Password: form.Password, Email: form.Email}
		dbHandler.Create(user)
		return EasyResponseJson{0, "注册成功"}
	}
	return EasyResponseJson{-1, "请填写必要信息"}
}

func UserLogin(form LoginForm) string {
	fmt.Println(form)
	return "Hello " + form.Username
}

