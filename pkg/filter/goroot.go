package filter

import (
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/awesome_libs/log"
	"go/build"
)

type GoRoot struct {
	Base
}

var GoRootFilter = GoRoot{
	Base{Name: "GoRootFilter"},
}

func (f GoRoot) Allow(targetObject Object) (allow bool) {
	if targetObject.Package == "" {
		allow = true
		debug(f, "", targetObject, allow)
		return
	}
	pkg, err := build.Import(targetObject.Package, "", build.FindOnly)
	if err != nil {
		awesome_error.CheckWarning(err)
		log.Logger.Warnf("pkg.Goroot=%v", pkg.Goroot)
	}
	allow = !pkg.Goroot
	debug(f, "pkg.Goroot", targetObject, allow)
	return
}
