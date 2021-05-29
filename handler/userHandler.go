package handler

import (
	"fmt"

	"MyBlog/db"

	"github.com/google/uuid"
	"github.com/kataras/golog"
)

type LoginForm struct {
	Username string `form:"username"`
	PassWord string `form:"password"`
}

type RegisterForm struct {
	Username  string `form:"username"`
	Email     string `form:"email"`
	Password  string `form:"password"`
	MobileNum string `form:"phoneNum"`
}

type EasyResponseJson struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func UserRegister(form RegisterForm) EasyResponseJson {
	fmt.Println(form)
	var middleUser db.User
	if form.Username != "" && form.Password != "" && form.Email != "" {
		dbHandler := db.GetDB()

		query := dbHandler.Where("name = ?", form.Username).Or("email = ?", form.Email).First(&middleUser)
		if query.RowsAffected != 0 {
			return EasyResponseJson{-1, "用户重复，请重新填写"}
		}

		var user = db.User{Name: form.Username, Password: form.Password, Email: form.Email, UID: uuid.New().String()}
		res := dbHandler.Create(&user)
		if res.Error != nil {
			golog.Error(res.Error)
			return EasyResponseJson{-1, "创建失败"}
		}

		return EasyResponseJson{0, "注册成功"}
	}

	return EasyResponseJson{-1, "请填写必要信息"}

}

func UserLogin(form LoginForm) EasyResponseJson {
	fmt.Println(form)
	var middleUser db.User
	if form.Username != "" && form.PassWord != "" {
		dbHandler := db.GetDB()

		queryName := dbHandler.Where("name = ?", form.Username).First(&middleUser)
		if queryName.RowsAffected == 0 {
			return EasyResponseJson{-1, "该用户名不存在！"}
		}
		queryPwd := dbHandler.Where("name = ? AND password = ? ", form.Username, form.PassWord).First(&middleUser)
		if queryPwd.RowsAffected == 0 {
			return EasyResponseJson{-1, "密码错误！"}
		}

		return EasyResponseJson{0, "登录成功"}

	}
	return EasyResponseJson{-1, "请填写必要信息"}

}
