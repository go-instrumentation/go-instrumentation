package filter

import (
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/awesome_libs/log"
	"go/build"
)

type GoRoot struct {
}

var GoRootFilter = GoRoot{}

func (f GoRoot) Match(pkgPath, _ string) (result bool) {
	pkg, err := build.Import(pkgPath, "", build.FindOnly)
	if err != nil {
		awesome_error.CheckWarning(err)
		log.Logger.Warnf("pkg.Goroot=%v", pkg.Goroot)
	}
	result = !pkg.Goroot
	if !result {
		log.Logger.Infof("pass %s because of pkg.Goroot", pkgPath)
	}
	return
}
