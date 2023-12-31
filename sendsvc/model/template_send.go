package model

import "time"

type TemplateSend struct {
	ID            int64     `db:"id,primary"`
	QueueID       string    `db:"queue_id"`
	Mobile        string    `db:"mobile"`
	NationCode    string    `db:"nationcode"`
	AppID         int       `db:"app_id"`
	TemplateType  string    `db:"t_template_type"`
	TemplateID    string    `db:"template_id"`
	TemplateParam string    `db:"template_param"`
	SendAt        string    `db:"send_at"`
	SendStatus    int       `db:"send_status"`
	SenderIP      string    `db:"sender_ip"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}

func (t TemplateSend) Table() string {
	return "t_template_send"
}
