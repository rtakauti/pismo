package main

type Account struct {
	Account_id                 int     `json:"id"`
	Available_credit_limit     float32 `json:"credit"`
	Available_withdrawal_limit float32 `json:"withdraw"`
}
