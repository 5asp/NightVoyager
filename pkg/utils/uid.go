package utils

import (
	"crypto/md5"
	"io"

	"github.com/rs/xid"
)

func CreateUid() string {
	guid := xid.New()
	h := md5.New()
	io.WriteString(h, guid.String())

	return string(h.Sum(nil))
}
