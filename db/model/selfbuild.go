package model

import (
	"github.com/go-instrumentation/go-instrumentation/pkg/util"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type SelfBuild struct {
	gorm.Model
	Pkg    string
	Binary string
}

var SchemaSelfBuild schema.Schema

func init() {
	awesome_error.CheckFatal(util.InitSchema(&SchemaSelfBuild, &SelfBuild{}))
}
