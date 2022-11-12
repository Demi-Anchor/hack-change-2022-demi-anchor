package handler

import "time"

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Donations struct {
	Name            string    `json:"name"`
	FromDate        time.Time `json:"from_date"`
	ToDate          time.Time `json:"to_date"`
	MoneyAmount     int64     `json:"money_amount"`
	DonationsAmount int       `json:"donations_amount"`
}
