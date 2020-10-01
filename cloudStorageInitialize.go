package main

import (
	"CloudStorage/conf"
	"CloudStorage/cs"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)


func newMysqlDsnString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		conf.CSConf.Mysql.User,
		conf.CSConf.Mysql.Password,
		conf.CSConf.Mysql.Host,
		conf.CSConf.Mysql.Port,
		conf.CSConf.Mysql.DBName,
	)
}

func MysqlInit() (err error) {
	return mysqlInit()
}

func mysqlInit() (err error) {
	mDsn := newMysqlDsnString()
	fmt.Println(mDsn)
	if cs.MySql, err = sqlx.Connect("mysql", mDsn); err != nil {
		return err
	}
	cs.MySql.SetMaxOpenConns(100)
	cs.MySql.SetMaxIdleConns(10)
	if err = cs.MySql.Ping(); err != nil {
		return err
	}
	return nil
}

func EtcdInit() {
	//TODO
}
