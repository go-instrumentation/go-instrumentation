package filter

import (
	"github.com/ssst0n3/awesome_libs/log"
)

func debug(f Filter, rule string, target Object, allow bool) {
	action := "deny"
	if allow {
		action = "allow"
	}
	log.Logger.Debugf("%+v %s %s because of %s", f.GetName(), action, target, rule)
}
