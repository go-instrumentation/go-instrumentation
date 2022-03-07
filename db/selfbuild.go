package db

import (
	"github.com/go-instrumentation/go-instrumentation/db/model"
	"github.com/ssst0n3/awesome_libs/awesome_error"
)

func ListSelfBuild() (selfBuildList []model.SelfBuild, err error) {
	err = DB.Model(&model.SelfBuild{}).Find(&selfBuildList).Error
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	return
}

func CreateSelfBuild(selfBuild model.SelfBuild) (err error) {
	lock.Lock()
	defer lock.Unlock()
	err = DB.Create(&selfBuild).Error
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	return
}
