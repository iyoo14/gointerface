package resource

import (
	"encoding/json"
	"github.com/iyoo14/gologger"
	"github.com/iyoo14/pqlap"
	"log"
	"os"
)

var logger *gologger.Logger
var con *pqlap.Db
var cfg Config

type Config struct {
	Dsn    string `json:"dsn"`
	Suffix string `json:suffix`
}

func SetCon(db *pqlap.Db) {
	con = db
}

func SetLogger(log *gologger.Logger) {
	logger = log
}

func GetCon() *pqlap.Db {
	return con
}

func GetLogger() *gologger.Logger {
	return logger
}

func SetConfig(jsonPath string) {
	f, err := os.Open(jsonPath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	err = json.NewDecoder(f).Decode(&cfg)
}

func GetConfig() Config {
	return cfg
}
