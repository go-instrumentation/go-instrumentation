package test_db

import (
	"github.com/go-instrumentation/go-instrumentation/db/model"
	"github.com/go-instrumentation/go-instrumentation/test/test_data"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/awesome_libs/awesome_reflect"
	"gorm.io/gorm"
)

func DeleteTable(db *gorm.DB, modelPtr interface{}) (err error) {
	err = db.Migrator().DropTable(modelPtr)
	awesome_error.CheckErr(err)
	return
}

func InitTable(db *gorm.DB, modelPtr interface{}, records interface{}) (err error) {
	awesome_reflect.MustPointer(modelPtr)
	awesome_reflect.MustPointer(records)
	err = DeleteTable(db, modelPtr)
	if err != nil {
		return
	}
	err = db.AutoMigrate(modelPtr)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	db.Create(records)
	return
}

func MakeTableEmpty(DB *gorm.DB, modelPtr interface{}) (err error) {
	awesome_reflect.MustPointer(modelPtr)
	err = DeleteTable(DB, modelPtr)
	if err != nil {
		return
	}
	err = DB.AutoMigrate(modelPtr)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	return
}

func InitSelfBuild(DB *gorm.DB) (err error) {
	return InitTable(DB, &model.SelfBuild{}, &test_data.SelfBuilds)
}
