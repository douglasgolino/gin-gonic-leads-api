package model

import "time"

type Lead struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	CreateDate time.Time `json:"create_date"`
}
