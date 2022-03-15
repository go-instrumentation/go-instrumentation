package db

import (
	"github.com/go-instrumentation/go-instrumentation/db/model"
	"github.com/ssst0n3/awesome_libs/log"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestPkgExistsInImportMap(t *testing.T) {
	sql := "/tmp/instrumentation.sqlite"
	assert.NoError(t, os.RemoveAll(sql))
	Init(sql)
	{
		exists, _, err := PkgExistsInImportMap("github.com/go-instrumentation/go-instrumentation/helper/jaeger")
		assert.NoError(t, err)
		assert.False(t, exists)
	}
	{
		assert.NoError(t, DB.Model(&model.Package{}).Create(&model.Package{PkgPath: "github.com/go-instrumentation/go-instrumentation/helper/jaeger"}).Error)
		exists, _, err := PkgExistsInImportMap("github.com/go-instrumentation/go-instrumentation/helper/jaeger")
		assert.NoError(t, err)
		assert.True(t, exists)
	}
}

func TestUpdatePkg(t *testing.T) {
	sql := "/tmp/instrumentation.sqlite"
	assert.NoError(t, os.RemoveAll(sql))
	Init(sql)
	log.Logger.Infof("%+v", DB.Config.ConnPool)
	pkgPath := "github.com/go-instrumentation/go-instrumentation/helper/jaeger"
	assert.NoError(t, DB.Model(&model.Package{}).Create(&model.Package{PkgPath: pkgPath}).Error)
	err := UpdatePkg(model.Package{
		PkgPath:        pkgPath,
		CfgPath:        "a",
		ImportPath:     "b",
		BinaryPath:     "c",
		CompileCommand: nil,
	})
	assert.NoError(t, err)
	pkg, err := FindPkgByPkgPath(pkgPath)
	assert.NoError(t, err)
	assert.Equal(t, pkg.CfgPath, "a")
}

func TestListPackages(t *testing.T) {
	sql := "/tmp/instrumentation.sqlite"
	assert.NoError(t, os.RemoveAll(sql))
	Init(sql)
	_, err := ListPackages()
	assert.NoError(t, err)
}
