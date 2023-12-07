package model

import "time"

type Account struct {
	ID        int       `db:"id,primary"`
	Account   string    `db:"account"`
	Password  string    `db:"password"`
	Status    int       `db:"status"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (t Account) Table() string {
	return "t_account"
}
