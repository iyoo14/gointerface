package repository

import (
	"github.com/iyoo14/pqlap"
)

func SelectTmpTable() bool {
	logger.Println("start selectTmpTable.")
	sql := `select id, name from tmp_table`
	Rows = con.SimpleQuery(sql)
	if con.Error() {
		logger.Println("prepare error.", sql)
		logger.Println(con.GetError())
		return false
	}
	return true
}

func InsertTestTable(record []interface{}) bool {
	logger.Println("start InsertTestTable.")
	sql := `insert into test_table
(id, name) values
($1, $2)`
	icon := pqlap.DbInstantConnection(cfg.Dsn)
	defer icon.Close()
	icon.Prepare(sql)
	if icon.Error() {
		logger.Println("prepare error.", sql)
		logger.Println(icon.GetError())
		logger.Println(record)
		return false
	}
	icon.Exec(record)
	if con.Error() {
		logger.Println("exec error.", sql)
		logger.Println(icon.GetError())
		logger.Println(record)
		return false
	}
	return true
}
