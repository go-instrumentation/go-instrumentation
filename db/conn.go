package db

import (
	"github.com/cloudquery/sqlite"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/awesome_libs/log"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"sync"
)

var DB *gorm.DB
var lock = &sync.Mutex{}
var Inited bool
var Dsn string

func InitGormDB(dsn string) (err error) {
	DB, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Discard,
	})
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	return
}

func Init(filepath string) {
	Inited = true
	Dsn = filepath
	log.Logger.Debugf(filepath)
	awesome_error.CheckFatal(InitGormDB(filepath))
	awesome_error.CheckFatal(Migrate())
}
