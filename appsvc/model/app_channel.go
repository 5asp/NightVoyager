package model

import (
	"encoding/json"
	"time"
)

type AppChannel struct {
	ID         int             `db:"id,primary"`
	AccountID  int             `db:"account_id"`
	AppID      int             `db:"app_id"`
	SendStatus int             `db:"send_status"`
	Reason     string          `db:"reason"`
	Extra      json.RawMessage `db:"extra"`
	CreatedAt  time.Time       `db:"created_at"`
}

func (t AppChannel) Table() string {
	return "t_app_channel"
}
