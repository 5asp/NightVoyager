package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io"

	"github.com/rs/xid"
)

func CreateUid() string {
	guid := xid.New()
	h := md5.New()
	io.WriteString(h, guid.String())

	return hex.EncodeToString(h.Sum(nil))
}
