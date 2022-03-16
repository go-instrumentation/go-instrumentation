package db

import (
	"errors"
	"fmt"
	"github.com/go-instrumentation/go-instrumentation/db/model"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"gorm.io/gorm"
)

func ListSelfBuild() (selfBuildList []model.SelfBuild, err error) {
	err = DB.Model(&model.SelfBuild{}).Find(&selfBuildList).Error
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	return
}

func SelfBuildExists(pkgPath string) (exists bool, selfBuild model.SelfBuild, err error) {
	selfBuild, err = FindSelfBuildByPkgPath(pkgPath)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			exists = false
			err = nil
			return
		}
		return
	}
	exists = selfBuild.Pkg != ""
	return
}

func FindSelfBuildByPkgPath(pkgPath string) (selfBuild model.SelfBuild, err error) {
	if pkgPath == "" {
		err = fmt.Errorf("empty pkgPath")
		awesome_error.CheckErr(err)
		return model.SelfBuild{}, err
	}
	err = DB.Model(&model.SelfBuild{}).Where(&model.SelfBuild{Pkg: pkgPath}).First(&selfBuild).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return
		}
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
