package model

import "time"

type Review struct {
	To      string    `json:"to"`
	Content string    `json:"content"`
	SendAt  time.Time `json:"send_at"`
}
