package model

import "time"

type AppInfo struct {
	ID        int       `db:"id,primary"`
	Secret    string    `db:"secret"`
	Status    int       `db:"status"`
	Remark    string    `db:"remark"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (t AppInfo) Table() string {
	return "t_sms_appinfo"
}
