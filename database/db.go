package main

import (
	"fmt"
	"log"

	"xorm.io/xorm"
)

type DB struct {
	engine *xorm.Engine
}

//ConnectDb open connection to db
func (d *DB) ConnectDb(sqlPath, dbName string) error {
	sqlConnStr := fmt.Sprintf("%s/%s?charset=utf8", sqlPath, dbName)
	engine, err := xorm.NewEngine("mysql", sqlConnStr)
	log.Print("Connect to: ", sqlConnStr)
	d.engine = engine
	return err
}
