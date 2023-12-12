package utils_test

import (
	"fmt"
	"strings"
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

func TestSplit(t *testing.T) {
	str := "【真北路】 早上好，绿洲中环中心."
	openBracket := "【"
	closeBracket := "】"

	startIndex := strings.Index(str, openBracket)
	endIndex := strings.Index(str, closeBracket)

	if startIndex == -1 || endIndex == -1 {
		// 未找到匹配的字符串
		fmt.Println("未找到匹配的字符串")
		return
	}

	content := str[startIndex+len(openBracket) : endIndex]
	remaining := strings.TrimSpace(str[endIndex+len(closeBracket):])
	fmt.Println("Content:", content)
	fmt.Println("Remaining:", remaining)

	result := []string{content, remaining}
	fmt.Println(result)
}
