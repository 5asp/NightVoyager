package model

import "time"

type AccountLog struct {
	ID        int       `db:"id,primary"`
	AccountID string    `db:"account_id"`
	IpAddr    string    `db:"ip_addr"`
	Device    string    `db:"device"`
	CreatedAt time.Time `db:"created_at"`
}

func (t AccountLog) Table() string {
	return "account_login"
}
