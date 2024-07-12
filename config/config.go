/*
 * @Description:
 * @Version: 2.0
 * @Autor: ABing
 * @Date: 2024-07-12 11:34:06
 * @LastEditors: lhl
 * @LastEditTime: 2024-07-12 11:48:54
 */
package config

import (
	"encoding/json"
	"log"

	"github.com/xelabs/go-mysqlstack/utils"
)

type Config struct {
	Username string   `json:"username"`
	Password string   `json:"password"`
	Readfile []string `json:"readfile"`
}

var GlobalConfig Config

func init() {

	// 解析 JSON 数据到 User 实例
	err := json.Unmarshal(utils.Read("config.json"), &GlobalConfig)

	if err != nil {
		log.Panicf(err.Error())
	}

	log.Println(GlobalConfig)
}
