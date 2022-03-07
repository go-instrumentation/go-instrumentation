package test_data

import (
	"github.com/go-instrumentation/go-instrumentation/db/model"
	"gorm.io/gorm"
)

var (
	SelfBuild1 = model.SelfBuild{
		Model: gorm.Model{
			ID: 1,
		},
		Pkg:    "SelfBuild1 pkg",
		Binary: "SelfBuild1 binary",
	}
	SelfBuild2 = model.SelfBuild{
		Model: gorm.Model{
			ID: 2,
		},
		Pkg:    "SelfBuild2 pkg",
		Binary: "SelfBuild2 binary",
	}
)

var SelfBuilds = []model.SelfBuild{
	SelfBuild1,
	SelfBuild2,
}
