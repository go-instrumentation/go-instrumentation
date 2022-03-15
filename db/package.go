package db

import (
	"errors"
	"fmt"
	"github.com/go-instrumentation/go-instrumentation/db/model"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"gorm.io/gorm"
)

func PkgExists(pkgPath string) (exists bool, pkg model.Package, err error) {
	pkg, err = FindPkgByPkgPath(pkgPath)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			exists = false
			err = nil
			return
		}
		return
	}
	exists = pkg.PkgPath != ""
	return
}

func MustPkgNotExists(pkgPath string) {
	exists, _, err := PkgExists(pkgPath)
	if exists {
		err = fmt.Errorf("pkg %s already exists", pkgPath)
	}
	awesome_error.CheckFatal(err)
}

func FindPkgByPkgPath(pkgPath string) (pkg model.Package, err error) {
	if pkgPath == "" {
		err = fmt.Errorf("empty pkgPath")
		awesome_error.CheckErr(err)
		return model.Package{}, err
	}
	err = DB.Model(&model.Package{}).Where(&model.Package{PkgPath: pkgPath}).First(&pkg).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return
		}
		awesome_error.CheckErr(err)
		return
	}
	return
}

func CreatePkg(pkg model.Package) (err error) {
	lock.Lock()
	defer lock.Unlock()
	err = DB.Create(&pkg).Error
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	return
}

func UpdatePkg(pkg model.Package) (err error) {
	lock.Lock()
	defer lock.Unlock()
	db := DB.Model(&model.Package{}).Where(&model.Package{PkgPath: pkg.PkgPath}).Updates(&pkg)
	err = db.Error
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	if db.RowsAffected == 0 {
		err = DB.Create(&pkg).Error
		if err != nil {
			awesome_error.CheckErr(err)
			return
		}
	}
	return
}

func ListPackages() (packages []model.Package, err error) {
	err = DB.Model(&model.Package{}).Find(&packages).Error
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	return
}
