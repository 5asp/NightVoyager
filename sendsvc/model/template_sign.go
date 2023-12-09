package model

import "time"

type TemplateSign struct {
	ID        int       `db:"id,primary"`
	Signature string    `db"signature"`
	IsDefault int       `db:"is_default"`
	AccountID int       `db:"account_id"`
	Status    int       `db:"status"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (t TemplateSign) Table() string {
	return "t_template_sign"
}
