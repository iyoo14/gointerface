package main

import (
	"database/sql"
	"github.com/iyoo14/gologger"
	"github.com/iyoo14/pqlap"
	"gointerface/repository"
	"gointerface/resource"
	"os"
	"path/filepath"
	"sync"
	"time"
)

const CMD string = "gointerface"

var logger *gologger.Logger
var baseDir string

func init() {
	exe, _ := os.Executable()
	exePath := filepath.Dir(exe)
	baseDir = filepath.Join(exePath, "..", "..", "..")
	jsonPath := filepath.Join(baseDir, "config", "config.json")
	resource.SetConfig(jsonPath)
	cfg := resource.GetConfig()
	logPath := filepath.Join(baseDir, "log")
	logger = gologger.NewLogger(logPath, CMD)
	con := pqlap.DbConnection(cfg.Dsn)
	if con.Error() {
		logger.Println("db connect error.")
		logger.Fatal(con.GetError())
	}
	resource.SetCon(con)
	resource.SetLogger(logger)
}

func main() {
	logger.Println("START ", CMD)
	repository.NewRepository()
	limit := make(chan struct{}, 3)
	var wg sync.WaitGroup
	repository.SelectTmpTable()
	targetRows := repository.Rows
	var id int
	for targetRows.Next() {
		id = id + 1
		logger.Println("loop ", id)
		wg.Add(1)
		if 1 != 1 {
			wg.Done()
			continue
		}
		go func(i int, tg *sql.Rows) {
			time.Sleep(time.Duration(i*500) * time.Microsecond)
			limit <- struct{}{}
			var id interface{}
			var name interface{}
			err := tg.Scan(&id, &name)
			defer wg.Done()
			defer func() {
				<-limit
			}()
			if err != nil {
				logger.Fatal("error rows:", err)
			}
			var record []interface{}
			record = append(record, id)
			record = append(record, name)
			repository.InsertTestTable(record)
			if err != nil {
				logger.Fatal("error insert:", err)
			}
			logger.Println("func ", i)
		}(id, targetRows)
		wg.Wait()
	}

}
