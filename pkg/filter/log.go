package filter

import (
	"github.com/ssst0n3/awesome_libs/log"
)

var LogSwitch = true

func check(f Filter, target Object) {
	if !LogSwitch {
		return
	}
	log.Logger.Debugf("%s start checking %s", f, target)
}

func debug(f Filter, rule string, target Object, allow bool) {
	if !LogSwitch {
		return
	}
	action := "deny"
	if allow {
		action = "allow"
	}
	log.Logger.Debugf("%s %s %s because of %s", f, action, target, rule)
}
