package models

import "time"

type Donation struct {
	StreamerID int       `json:"streamer_id"`
	Author     string    `json:"author"`
	Money      int64     `json:"money"`
	Comment    string    `json:"comment"`
	Time       time.Time `json:"time"`
}

type Period struct {
	FirstDate time.Time `json:"first_date"`
	LastDate  time.Time `json:"last_date"`
}

type DailyDonation struct {
	Sum   int64     `json:"sum" db:"sum"`
	Count int       `json:"count" db:"count"`
	Date  time.Time `json:"date" db:"date"`
}
