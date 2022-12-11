package database

import (
	"flag"
	"marllef/beautiful-api/configs"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	dbLog "gorm.io/gorm/logger"
)

var dbDialect = flag.String("db-dialect", "sqlite", "set database dialect, sqlite by default.")

func NewDB() (container *gorm.DB, err error) {
	container, err = gorm.Open(sqlite.Open(configs.Env.DatabaseUrl), &gorm.Config{
		Logger: dbLog.Default.LogMode(dbLog.Silent),
	})

	if err != nil {
		return nil, err
	}

	
	return container, nil
}
