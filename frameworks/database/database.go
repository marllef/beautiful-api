package database

import (
	"flag"
	"marllef/beautiful-api/configs"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	dbLog "gorm.io/gorm/logger"
)

var dbDialect = flag.String("db-dialect", "sqlite", "set database dialect, sqlite by default.")

type Database interface {}

type database struct {
	container *gorm.DB
	logLevel  int
	Database
}

func NewDatabase() (container *gorm.DB, err error) {
	
	if err = configs.LoadEnvs(); err != nil {
		return nil, err
	}

	container, err = gorm.Open(sqlite.Open(configs.Env.DatabaseUrl), &gorm.Config{
		Logger: dbLog.Default.LogMode(dbLog.Silent),
	})

	if err != nil {
		return nil, err
	}

	return container, nil
}
