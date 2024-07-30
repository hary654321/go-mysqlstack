/*
 * @Description:
 * @Version: 2.0
 * @Autor: ABing
 * @Date: 2024-07-10 10:09:31
 * @LastEditors: lhl
 * @LastEditTime: 2024-07-30 17:59:04
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
	"os"
	"os/signal"
	"syscall"

	"github.com/xelabs/go-mysqlstack/driver"
	"github.com/xelabs/go-mysqlstack/jsonlog"
	"github.com/xelabs/go-mysqlstack/mock"
	"github.com/xelabs/go-mysqlstack/utils"
	"github.com/xelabs/go-mysqlstack/xlog"
)

func main() {

	jsonlog.Init()

	log := xlog.NewStdLog(xlog.Level(xlog.INFO))
	th := driver.NewTestHandler(log)

	result1 := mock.TableRes()
	th.AddQuery("SELECT * FROM MOCK", result1)

	result2 := mock.ConnectionID()
	th.AddQuery("SELECT CONNECTION_ID()", result2)

	result3 := mock.ShowVal("showval.json")
	th.AddQuery("SHOW VARIABLES", result3)

	result4 := mock.ShowVal("showstatus.json")
	th.AddQuery("SHOW STATUS", result4)
	th.AddQuery("SHOW  STATUS", result4)

	result5 := mock.ShowDataBase("showDatabase.json")
	th.AddQuery("SHOW DATABASES", result5)

	mysqld, err := driver.MockMysqlServerWithPort(log, utils.GetHpPort(), th)
	if err != nil {
		log.Panic("mysqld.start.error:%+v", err)
	}
	defer mysqld.Close()
	log.Info("mysqld.server.start.address[%v]", mysqld.Addr())

	// Handle SIGINT and SIGTERM.
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
}
