package filter

import (
	"github.com/sirupsen/logrus"
	"github.com/ssst0n3/awesome_libs/log"
	"testing"
)

func Test_debug(t *testing.T) {
	log.Logger.Level = logrus.DebugLevel
	debug(RuleDenyGoInstrumentationFamily, "rule", Object{Package: "pkg"}, false)
}
