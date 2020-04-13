package main

import (
	"./db"
	"./handler"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)

func main() {
	db.DBInit()
	app := iris.New()
	app.Logger().SetLevel("debug")
	// Optionally, add two built'n handlers
	// that can recover from any http-relative panics
	// and log the requests to the terminal.
	app.Use(recover.New())
	app.Use(logger.New())

	// Method:   GET
	// Resource: http://localhost:8080
	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("<h1>Welcome</h1>")
		fmt.Println()
	})

	// same as app.Handle("GET", "/ping", [...])
	// Method:   GET
	// Resource: http://localhost:8080/ping
	app.Get("/ping/{id:uint64}", func(ctx iris.Context) {
		ctx.WriteString("pong")
		id := ctx.Params().GetUint64Default("id", 0)
		fmt.Println(id)
	})

	// Method:   GET
	// Resource: http://localhost:8080/hello
	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello Iris!"})
	})

	hero.Register(func (ctx iris.Context) (form handler.LoginForm){
		ctx.ReadForm(&form)
		fmt.Println("Login: ", form)
		return
	})
	loginHandler := hero.Handler(handler.UserLogin)
	app.Post("/login", loginHandler)

	hero.Register(func (ctx iris.Context) (form handler.RegisterForm){
		ctx.ReadForm(&form)
		fmt.Println("Register: ", form)
		return
	})
	registerHandler := hero.Handler(handler.UserRegister)
	app.Post("/register", registerHandler)

	// http://localhost:8080
	// http://localhost:8080/ping
	// http://localhost:8080/hello

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
