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

	transactionRepo := NewTransactionRepository(DBConn(conf))
	transactionService := NewTransactionService(transactionRepo)

	app.Get("/v1/accounts", func(ctx iris.Context) {
		ctx.JSON(accountService.GetAll())
	})

	app.Get("/v1/accounts/limits", func(ctx iris.Context) {
		ctx.JSON(accountService.GetAll())
	})

	app.Get("/v1/accounts/{id:int}", func(ctx iris.Context) {
		id, _ := ctx.Params().GetInt("id")
		account := accountService.GetById(id)
		if account.Account_id != 0 {
			ctx.JSON(account)
		}

		ctx.StatusCode(iris.StatusNotFound)
	})

	app.Delete("/v1/accounts/{id:int}", func(ctx iris.Context) {
		id, _ := ctx.Params().GetInt("id")
		exists := accountService.Delete(id)
		if !exists {
			ctx.StatusCode(iris.StatusNotFound)
		} else {
			ctx.StatusCode(iris.StatusNoContent)
		}

		return
	})

	app.Post("/v1/accounts", func(ctx iris.Context) {
		var account Account
		ctx.ReadJSON(&account)

		inserterAccount, err := accountService.Insert(account)
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)

			return
		}
		ctx.JSON(inserterAccount)
	})

	app.Put("/v1/accounts/{id:int}", func(ctx iris.Context) {
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

	app.Patch("/v1/accounts/{id:int}", func(ctx iris.Context) {
		var account Account
		ctx.ReadJSON(&account)
		id, _ := ctx.Params().GetInt("id")
		cont := accountService.GetById(id)
		if cont.Account_id == 0 {
			ctx.StatusCode(iris.StatusNotFound)
		}

		updatedAccount, err := accountService.Patch(id, account)
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)

			return
		}
		ctx.JSON(updatedAccount)
	})

	app.Post("/v1/transactions", func(ctx iris.Context) {
		var transaction Transaction
		ctx.ReadJSON(&transaction)

		inserterTransaction, err := transactionService.Insert(transaction)
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)

			return
		}
		ctx.JSON(inserterTransaction)
	})

	app.Run(
		iris.Addr(":80"),
		iris.WithOptimizations,
	)
}
