package model

import "time"

type App struct {
	ID        int       `db:"id,primary"`
	Name      string    `db:"name"`
	Secret    string    `db:"secret"`
	Status    int       `db:"status"`
	Remark    string    `db:"remark"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	AccountID int       `db:"account_id"`
}

func (t App) Table() string {
	return "t_app"
}
