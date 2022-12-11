package database

import (
	"flag"
	"marllef/beautiful-api/configs"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	dbLog "gorm.io/gorm/logger"
)

var dbDialect = flag.String("db-dialect", "sqlite", "set database dialect, sqlite by default.")
var container *gorm.DB

func NewDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(configs.Env.DatabaseUrl), &gorm.Config{
		Logger: dbLog.Default.LogMode(dbLog.Silent),
	})

	if err != nil {
		return nil, err
	}

	container = db
	return container, nil
}
