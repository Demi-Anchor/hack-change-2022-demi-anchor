package models

import "time"

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Payment struct {
	Name  string `json:"name"`
	Money int64  `json:"money"`
}

type Period struct {
	FirstDate time.Time `json:"first_date"`
	LastDate  time.Time `json:"last_date"`
}
