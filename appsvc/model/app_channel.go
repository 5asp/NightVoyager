package model

import (
	"time"
)

type AppChannel struct {
	ID            int       `db:"id,primary"`
	ChannelName   string    `db:"channel_name"`
	ChannelType   string    `db:"channel_type"`
	ChannelKey    string    `db:"channel_key"`
	ChannelDomain string    `db:"channel_domain"`
	ExtProperties string    `db:"ext_properties"`
	IsDefault     int       `db:"is_default"`
	SendOrder     int       `db:"send_order"`
	Quota         int       `db:"quota"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}

func (t AppChannel) Table() string {
	return "t_app_channel"
}
