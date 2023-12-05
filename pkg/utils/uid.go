package utils

import "github.com/rs/xid"

func CreateUid() string {
	guid := xid.New()
	return guid.String()
}
