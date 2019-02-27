package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type TransactionRepository interface {
	GetAll() (results []Transaction)
	Get(transaction Transaction) (results []Transaction)
	GetById(id int64) (result Transaction)
	Insert(transaction Transaction) (insertedTransaction Transaction, err error)
	Update(id int64, transaction Transaction) (updatedTransaction Transaction, err error)
	Delete(id int64) (deleted bool)
}

type dbTransactionRepository struct {
	Conn *sql.DB
}

func NewTransactionRepository(Conn *sql.DB) *dbTransactionRepository {
	return &dbTransactionRepository{
		Conn: Conn,
	}
}

func (repository *dbTransactionRepository) GetAll() (results []Transaction) {
	selDB, err := repository.Conn.Query("SELECT * FROM transactions")

	if err != nil {
		panic(err.Error())
	}
	for selDB.Next() {
		transaction := Transaction{}
		err = selDB.Scan(&transaction.Transaction_id, &transaction.Account_id, &transaction.Operation_type_id, &transaction.Amount, &transaction.Balance, &transaction.Created_at, &transaction.Due_at )
		if err != nil {
			panic(err.Error())
		}

		results = append(results, transaction)
	}

	return results
}

func (repository *dbTransactionRepository) Get(transaction Transaction) (results []Transaction) {

	return results
}

func (repository *dbTransactionRepository) GetById(id int64) (result Transaction) {
	stmt, err := repository.Conn.Prepare("SELECT * FROM transactions where Transaction_id = ?")
	if err != nil {
		panic(err.Error())
	}

	res, err := stmt.Query(id)
	for res.Next() {
		transaction := Transaction{}
		err = res.Scan(&transaction.Transaction_id, &transaction.Account_id, &transaction.Operation_type_id, &transaction.Amount, &transaction.Balance, &transaction.Created_at, &transaction.Due_at )
		if err != nil {
			panic(err.Error())
		}

		result = transaction
	}

	defer res.Close()

	return
}

func (repository *dbTransactionRepository) Insert(transaction Transaction) (insertedTransaction Transaction, err error) {
	stmt, err := repository.Conn.Prepare("INSERT INTO transactions (Account_id, Operation_type_id, Amount, Balance) VALUES(?, ?, ?, ?)")
	transaction.Amount = transaction.Amount * -1
	transaction.Balance = transaction.Balance * -1
	res, err := stmt.Exec(transaction.Account_id, transaction.Operation_type_id, transaction.Amount, transaction.Balance)
	if err != nil {
		panic(err.Error())
	}

	id, err := res.LastInsertId()
	if err != nil {
		panic(err.Error())
	}

	transaction.Transaction_id = int64(id)
	insertedTransaction = transaction

	return insertedTransaction, err
}

func (repository *dbTransactionRepository) Update(id int64, transaction Transaction) (updatedTransaction Transaction, err error) {
	stmt, err := repository.Conn.Prepare("UPDATE transactions SET Account_id = ?, Operation_type_id = ?, Amount = ?, Balance = ?, Due_at = ? WHERE Transaction_id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(transaction.Account_id, transaction.Operation_type_id, transaction.Amount, transaction.Balance, transaction.Due_at, id)
	if err != nil {
		panic(err.Error())
	}

	updatedTransaction = transaction
	updatedTransaction.Transaction_id = id

	return updatedTransaction, err
}

func (repository *dbTransactionRepository) Delete(id int64) (deleted bool) {
	stmt, err := repository.Conn.Prepare("DELETE FROM transactions where Transaction_id = ?")

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
