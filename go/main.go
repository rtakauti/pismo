package main

import (
	"github.com/kataras/iris"
	"os"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	conf := &DatabaseConf{
		DBDriver: "mysql",
		DBUser:   os.Getenv("DB_USER"),
		DBPass:   os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}

	repo := NewAccountRepository(DBConn(conf))
	accountService := NewAccountService(repo)

	app.Get("/accounts", func(ctx iris.Context) {
		ctx.JSON(accountService.GetAll())
	})

	app.Get("/accounts/{id:int}", func(ctx iris.Context) {
		id, _ := ctx.Params().GetInt("id")
		account := accountService.GetById(id)
		if account.Account_id != 0 {
			ctx.JSON(account)
		}

		ctx.StatusCode(iris.StatusNotFound)
	})

	app.Delete("/accounts/{id:int}", func(ctx iris.Context) {
		id, _ := ctx.Params().GetInt("id")
		exists := accountService.Delete(id)
		if !exists {
			ctx.StatusCode(iris.StatusNotFound)
		} else {
			ctx.StatusCode(iris.StatusNoContent)
		}

		return
	})

	app.Post("/accounts", func(ctx iris.Context) {
		var account Account
		ctx.ReadJSON(&account)

		inserterAccount, err := accountService.Insert(account)
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)

			return
		}
		ctx.JSON(inserterAccount)
	})

	app.Put("/accounts/{id:int}", func(ctx iris.Context) {
		var account Account
		ctx.ReadJSON(&account)
		id, _ := ctx.Params().GetInt("id")
		cont := accountService.GetById(id)
		if cont.Account_id == 0 {
			ctx.StatusCode(iris.StatusNotFound)
		}

		updatedAccount, err := accountService.Update(id, account)
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)

			return
		}
		ctx.JSON(updatedAccount)
	})

	app.Run(
		iris.Addr(":80"),
		iris.WithOptimizations,
	)
}
