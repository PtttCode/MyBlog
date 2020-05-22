package main

import (
	"./db"
	"./handler"
	"fmt"
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)

func h(ctx iris.Context){
	name, passwd, _ := ctx.Request().BasicAuth()
	ctx.JSON(iris.Map{
		"path": ctx.Path(),
		"name": name,
		"passwd": passwd,
	})
}



func main() {
	db.DBInit()
	app := iris.New()
	app.Logger().SetLevel("debug")

	app.Use(recover.New())
	app.Use(logger.New())

	app.Get("/ping/{id:uint64}", func(ctx iris.Context) {
		ctx.WriteString("pong")
		id := ctx.Params().GetUint64Default("id", 0)
		fmt.Println(id)
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

	hero.Register(func (ctx iris.Context) (jsondata handler.DocumentInput){
		ctx.ReadJSON(&jsondata)
		golog.Info(jsondata)
		return
	})
	docInsertHandler := hero.Handler(handler.InsertHandler)
	app.Put("/docs/insert", docInsertHandler)

	hero.Register(func (ctx iris.Context) (jsondata handler.DocumentInput){
		ctx.ReadJSON(&jsondata)
		golog.Info(jsondata)
		return
	})
	docGetHandler := hero.Handler(handler.GetDocHandler)
	app.Post("/docs/get", docGetHandler)
	//app.PartyFunc("/docs", func(r iris.Party) {
	//	//r.Use(middleware.BasicAuth)
	//	r.Post("/get_docs", hero.Handler(handler.GetDocHandler))
	//	r.Put("/insert_doc", hero.Handler(handler.InsertHandler))
	//})


	// http://localhost:8080
	// http://localhost:8080/ping
	// http://localhost:8080/hello
	//app.Listen("0.0.0.0:5555")

	//html := iris.HTML("./static/html", ".html")
	//html.Layout("index.html")
	//html.Reload(true)
	//app.RegisterView(html)
	app.HandleDir("/static", "./static")
	app.Layout("static/index.html")

	//app.Get("/", func(ctx iris.Context) {
	//	ctx.ViewData("message", "Welcome!")
	//	ctx.View("index.html")
	//
	//})

	app.Run(
		iris.Addr(":5555"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
		)
}
