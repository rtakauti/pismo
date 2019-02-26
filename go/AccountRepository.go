package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type AccountRepository interface {
	GetAll() (results []Account)
	Get(account Account) (results []Account)
	GetById(id int) (result Account)
	Insert(account Account) (insertedAccount Account, err error)
	Update(id int, account Account) (updatedAccount Account, err error)
	Delete(id int) (deleted bool)
}

type dbAccountRepository struct {
	Conn *sql.DB
}

func NewAccountRepository(Conn *sql.DB) *dbAccountRepository {
	return &dbAccountRepository{
		Conn: Conn,
	}
}

func (repository *dbAccountRepository) GetAll() (results []Account) {
	selDB, err := repository.Conn.Query("SELECT * FROM accounts")

	if err != nil {
		panic(err.Error())
	}
	for selDB.Next() {
		account := Account{}
		err = selDB.Scan(&account.Account_id, &account.Available_credit_limit, &account.Available_withdrawal_limit)
		if err != nil {
			panic(err.Error())
		}

		results = append(results, account)
	}

	return results
}

func (repository *dbAccountRepository) Get(account Account) (results []Account) {

	return results
}

func (repository *dbAccountRepository) GetById(id int) (result Account) {
	stmt, err := repository.Conn.Prepare("SELECT * FROM accounts where Account_id = ?")
	if err != nil {
		panic(err.Error())
	}

	res, err := stmt.Query(id)
	for res.Next() {
		account := Account{}
		err = res.Scan(&account.Account_id, &account.Available_credit_limit, &account.Available_withdrawal_limit)
		if err != nil {
			panic(err.Error())
		}

		result = account
	}

	defer res.Close()

	return
}

func (repository *dbAccountRepository) Insert(account Account) (insertedAccount Account, err error) {
	stmt, err := repository.Conn.Prepare("INSERT INTO accounts (Available_credit_limit, Available_withdrawal_limit) VALUES(?, ?)")
	res, err := stmt.Exec(account.Available_credit_limit, account.Available_withdrawal_limit)
	if err != nil {
		panic(err.Error())
	}

	id, err := res.LastInsertId()
	if err != nil {
		panic(err.Error())
	}

	account.Account_id = int(id)
	insertedAccount = account

	return insertedAccount, err
}

func (repository *dbAccountRepository) Update(id int, account Account) (updatedAccount Account, err error) {
	stmt, err := repository.Conn.Prepare("UPDATE accounts SET Available_credit_limit = ?, Available_withdrawal_limit = ? WHERE Account_id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(account.Available_credit_limit, account.Available_withdrawal_limit, id)
	if err != nil {
		panic(err.Error())
	}

	updatedAccount = account
	updatedAccount.Account_id = id

	return updatedAccount, err
}

func (repository *dbAccountRepository) Delete(id int) (deleted bool) {
	stmt, err := repository.Conn.Prepare("DELETE FROM accounts where Account_id = ?")

	if err != nil {
		panic(err.Error())
	}
	res, err := stmt.Exec(id)

	if err != nil {
		panic(err.Error())
	}

	rows, err := res.RowsAffected()

	if err != nil {
		panic(err.Error())
	}

	deleted = false
	if rows > 0 {
		deleted = true
	}

	return
}
