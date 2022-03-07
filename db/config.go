package db

import (
	"errors"
	"github.com/go-instrumentation/go-instrumentation/db/model"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"gorm.io/gorm"
)

func IsPrebuilt() (prebuilt bool, err error) {
	var config model.Config
	err = DB.Model(&model.Config{}).Where(&model.Config{Key: "prebuilt"}).First(&config).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = nil
			return
		}
		awesome_error.CheckErr(err)
		return
	}
	prebuilt = true
	return
}

func Prebuilt() (err error) {
	lock.Lock()
	defer lock.Unlock()
	err = DB.Model(&model.Config{}).Create(&model.Config{Key: "prebuilt", Value: "yes"}).Error
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	return
}
