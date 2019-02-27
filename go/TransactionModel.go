package main

import "time"

type Transaction struct {
	Transaction_id    int64     `json:"transaction_id"`
	Account_id        int       `json:"account_id"`
	Operation_type_id int       `json:"operation_type_id"`
	Amount            float32   `json:"amount"`
	Balance           float32   `json:"balance"`
	Created_at        time.Time `json:"created_at"`
	Due_at            time.Time `json:"due_at"`
}
