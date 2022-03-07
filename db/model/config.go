package model

import (
	"github.com/go-instrumentation/go-instrumentation/pkg/util"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Config struct {
	gorm.Model
	Key   string
	Value string
}

var SchemaConfig schema.Schema

func init() {
	awesome_error.CheckFatal(util.InitSchema(&SchemaConfig, &Config{}))
}
