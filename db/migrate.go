package db

import (
	"github.com/go-instrumentation/go-instrumentation/db/model"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/awesome_libs/awesome_reflect"
)

func migrate(m interface{}) (err error) {
	awesome_reflect.MustPointer(m)
	if DB.Migrator().HasTable(m) {
		return
		//err = DB.Migrator().AutoMigrate(m)
		//if err != nil {
		//	awesome_error.CheckErr(err)
		//	return
		//}
	} else {
		err = DB.Migrator().CreateTable(m)
		if err != nil {
			awesome_error.CheckErr(err)
			return
		}
	}
	return
}

func MigrateTables(models ...interface{}) (err error) {
	for _, m := range models {
		err = migrate(m)
		if err != nil {
			return err
		}
	}
	return
}

func Migrate() (err error) {
	return MigrateTables(&model.SelfBuild{}, &model.Package{}, &model.Config{})
}
