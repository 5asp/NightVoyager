package utils_test

import (
	"fmt"
	"testing"

	"github.com/aheadIV/NightVoyager/pkg/utils"
)

func TestCreateID(t *testing.T) {
	fmt.Println(utils.CreateID())
}

func TestEncode(t *testing.T) {
	code := utils.Encode(9999900009999900000)
	fmt.Println(len(utils.BASE))
	fmt.Println(code)
	fmt.Println(utils.Decode(code))
}
