package filter

import (
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/awesome_libs/log"
	"go/build"
)

type GoRoot struct {
}

var GoRootFilter = GoRoot{}

func (f GoRoot) Allow(targetObject Object) (allow bool) {
	if targetObject.Package == "" {
		allow = true
		return
	}
	pkg, err := build.Import(targetObject.Package, "", build.FindOnly)
	if err != nil {
		awesome_error.CheckWarning(err)
		log.Logger.Warnf("pkg.Goroot=%v", pkg.Goroot)
	}
	allow = !pkg.Goroot
	if !allow {
		log.Logger.Infof("pass %s because of pkg.Goroot", targetObject.Package)
	}
	return
}
