/*
 * @Description:
 * @Version: 2.0
 * @Autor: ABing
 * @Date: 2024-07-12 11:42:56
 * @LastEditors: lhl
 * @LastEditTime: 2024-07-12 11:44:04
 */
package utils

import (
	"io"
	"os"
)

func Read(filename string) []byte {
	// 打开 JSON 文件
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 读取文件内容
	jsonBytes, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	return jsonBytes
}
