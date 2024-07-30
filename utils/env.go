package utils

import (
	"log"
	"os"
	"strconv"
)

func GetHpName() string {

	home := os.Getenv("HP_NAME")
	if home == "" {
		return "mysql"
	}

	return home

}

func GetHpPort() int {

	home := os.Getenv("HP_PORT")
	if home == "" {
		return 3306
	}

	num, _ := strconv.Atoi(home)
	return num

}

func GetHpLogPath() string {

	home := os.Getenv("HP_LOG_PATH")

	log.Println("lllllllll", home)

	if home == "" {
		return "./log/mysql.json"
	}

	return home

}

func GetLoginName() string {

	home := os.Getenv("LOGIN_NAME")
	if home == "" {
		return "root"
	}

	return home

}

func GetLoginPwd() string {

	home := os.Getenv("LOGIN_PWD")
	if home == "" {
		return "root"
	}

	return home

}
