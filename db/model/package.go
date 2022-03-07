package model

import (
	"github.com/go-instrumentation/go-instrumentation/pkg/util"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Package struct {
	gorm.Model
	PkgPath        string
	CfgPath        string
	ImportPath     string
	BinaryPath     string
	Fingerprint    string
	CompileCommand []byte
}

var SchemaPackage schema.Schema

func init() {
	awesome_error.CheckFatal(util.InitSchema(&SchemaPackage, &Package{}))
}
