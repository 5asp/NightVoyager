package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"

	"github.com/rs/xid"
)

func CreateUid() string {
	guid := xid.New()
	h := md5.New()
	io.WriteString(h, guid.String())

	return hex.EncodeToString(h.Sum(nil))
}

func CreateID() string {
	guid := xid.New()
	fmt.Println(guid.Pid())
	fmt.Println(guid.Counter())
	fmt.Println(string(guid.Machine()))
	return guid.String()
}
