package go_instrumentation

import (
	"github.com/sirupsen/logrus"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/awesome_libs/log"
	"os"
	"time"
)

var WorkId string

func init() {
	log.Logger.Level = logrus.DebugLevel
	err := log.Output2File("/tmp/install_instrumentation")
	if err != nil {
		os.Exit(-1)
	}
	WorkId = time.Now().String()
	log.Logger.Infof("======%s START====", WorkId)
}

func Finish() {
	defer func() {
		awesome_error.CheckErr(log.CloseFile())
	}()
	log.Logger.Infof("=====%s FINISHED====\n\n", WorkId)
}
