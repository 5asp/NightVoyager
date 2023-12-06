package model

import "time"

type AccountLog struct {
	ID        int       `db:"id,primary"`
	Account   string    `db:"account"`
	Password  string    `db:"password"`
	Status    int       `db:"status"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (t AccountLog) Table() string {
	return "account_login"
}
