/*
 * @Description:
 * @Version: 2.0
 * @Autor: ABing
 * @Date: 2024-07-12 15:00:31
 * @LastEditors: lhl
 * @LastEditTime: 2024-07-12 17:38:09
 */
/*
 * go-mysqlstack
 * xelabs.org
 *
 * Copyright (c) XeLabs
 * GPL License
 *
 */

package utils

import (
	"fmt"
	"log"
	"testing"
)

func TestCache(t *testing.T) {
	k := "1.1.2"
	SetItem(k, 1)

	log.Println(GetItem(k))
	log.Println(GetItem("a"))
}

func TestOx(t *testing.T) {
	var TRBULAR = []byte{
		0xfb,
		0x44, 0x3a, 0x2f, 0x69, 0x6f, 0x63, 0x2e, 0x74, 0x78, 0x74,
	}

	// 将字节切片转换为字符串
	str := string(TRBULAR)

	// 输出转换后的字符串
	fmt.Println(str)
}

func TestPayload(t *testing.T) {

	// 输出转换后的字符串
	fmt.Println(GetPayload("D:\\ioct.txt"))

	fmt.Println(string(GetPayload("D:\\ioct.txt")))
}

func TestGetUserName(t *testing.T) {

	// 输出转换后的字符串
	fmt.Println(GetUserName(`
 2/27/2024 15:9:4 - PFRO Error: \??\C:\Program Files\Google\Chrome\Temp\scoped_dir12712_18953496\old_chrome.exe, |delete operation|, 0xc000003a
2/27/2024 15:9:4 - PFRO Error: \??\C:\Program Files\Google\Chrome\Temp\scoped_dir12712_18953496, |delete operation|, 0xc0000034
2/27/2024 15:9:4 - PFRO Error: \??\C:\Program Files\Google\Chrome\Temp, |delete operation|, 0xc0000101
2/27/2024 15:9:4 - PFRO Error: \??\C:\Users\zw\AppData\Local\Temp\iu-14D2N.tmp\_unins.tmp, |delete operation|, 0xc000003a
2/27/2024 15:9:4 - PFRO Error: \??\C:\Users\zw\AppData\Local\Temp\iu-14D2N.tmp, |delete operation|, 0xc0000034
2/27/2024 15:9:4 - PFRO Error: \??\C:\Users\zw\AppData\Local\Temp\{9A9905B7-8E17-4ef4-B3F2-C4CA89D18195}.tmp, |delete operation|, 0xc0000034
2/27/2024 15:9:4 - PFRO Error: \??\C:\Users\zw\AppData\Local\Temp\{D6F8EBE8-E015-4210-AFD7-70C627CA28AA}.tmp\MiniUI.dll, |delete operation|, 0xc000003a
2/27/2024 15:9:4 - PFRO Error: \??\C:\Users\zw\AppData\Local\Temp\{D6F8EBE8-E015-4210-AFD7-70C627CA28AA}.tmp, |delete operation|, 0xc0000034
2/27/2024 15:9:4 - PFRO Error: \??\C:\Users\zw\AppData\Local\Temp\offline1312.cab, |delete operation|, 0xc0000034
2/27/2024 15:9:4 - PFRO Error: \??\C:\Program Files (x86)\360\360Safe\update\~TH4084.cab, |delete operation|, 0xc0000034
2/27/2024 15:9:4 - PFRO Error: \??\C:\WINDOWS\SysWOW64\drivers\360boost.sys.429, |delete operation|, 0xc0000034
2/27/2024 15:9:4 - PFRO Error: \??\C:\WINDOWS\SysWOW64\drivers\360boost.sys.315, |delete operation|, 0xc0000034
2/27/2024 15:9:4 - PFRO Error: \??\C:\Users\zw\AppData\Local\Temp\~RSE667.tmp, |delete operation|, 0xc0000034
2/27/2024 15:9:4 - PFRO Error: \??\C:\Users\zw\AppData\Local\Temp\nshE1E4.tmp\nsProcess.dll, |delete operation|, 0xc000003a
2/27/2024 15:9:4 - PFRO Error: \??\C:\Users\zw\AppData\Local\Temp\nshE1E4.tmp\, |delete operation|, 0xc0000034
2/27/2024 15:9:4 - PFRO Error: \??\C:\Users\zw\AppData\Local\Temp\~RSB6AC.tmp, |delete operation|, 0xc0000034
2/27/2024 15:9:4 - PFRO Error: \??\C:\WINDOWS\SysWOW64\drivers\360boost.sys.0	`))

}

func TestGetGetWechatId(t *testing.T) {

	filecontent := Read("D:\\wz\\WeChat Files\\All Users\\config\\config.data")

	log.Println(string(filecontent))
	// 输出转换后的字符串
	log.Println("GetWechatId----", GetWechatId(string(filecontent)))

}
