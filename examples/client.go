/*
 * @Description:
 * @Version: 2.0
 * @Autor: ABing
 * @Date: 2024-07-10 10:09:31
 * @LastEditors: lhl
 * @LastEditTime: 2024-07-10 14:36:41
 */
/*
 * go-mysqlstack
 * xelabs.org
 *
 * Copyright (c) XeLabs
 * GPL License
 *
 */

package main

import (
	"fmt"

	"github.com/xelabs/go-mysqlstack/driver"
	"github.com/xelabs/go-mysqlstack/xlog"
)

func main() {
	log := xlog.NewStdLog(xlog.Level(xlog.INFO))
	address := fmt.Sprintf(":4407")
	client, err := driver.NewConn("root", "root", address, "", "")
	if err != nil {
		log.Panic("client.new.connection.error:%+v", err)
	}
	defer client.Close()

	qr, err := client.FetchAll("SELECT * FROM MOCK", -1)
	if err != nil {
		log.Panic("client.query.error:%+v", err)
	}
	log.Info("results:[%+v]", qr.Rows)
}
