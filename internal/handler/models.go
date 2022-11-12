package handler

import "time"

type Donations struct {
	Name            string    `json:"name"`
	FromDate        time.Time `json:"from_date"`
	ToDate          time.Time `json:"to_date"`
	MoneyAmount     int64     `json:"money_amount"`
	DonationsAmount int       `json:"donations_amount"`
}
