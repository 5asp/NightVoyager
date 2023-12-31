package model

import "time"

type Template struct {
	ID         int64     `db:"id,primary"`
	QueueID    string    `db:"queue_id"`
	TplContent string    `db:"tpl_content"`
	TplType    string    `db:"tpl_type"`
	SignID     int       `db:"sign_id"`
	AccountID  int       `db:"account_id"`
	Status     int       `db:"status"`
	Remark     string    `db:"remark"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

func (t Template) Table() string {
	return "t_template"
}
