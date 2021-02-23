package repository

import (
	"database/sql"
	"github.com/iyoo14/gologger"
	"github.com/iyoo14/pqlap"
	"gointerface/resource"
)

var logger *gologger.Logger
var con *pqlap.Db
var Rows *sql.Rows

func NewRepository() {
	con = resource.GetCon()
	logger = resource.GetLogger()
}
