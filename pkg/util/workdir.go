package util

import (
	"fmt"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"regexp"
)

func FindGoBuildWorkDir(output []byte) (workDir string, err error) {
	r, err := regexp.Compile("WORK=(/.*/go-build[0-9]+)")
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	match := r.FindSubmatch(output)
	if len(match) != 2 {
		err = fmt.Errorf("regex match failed: %s, the origin string is %+v", match, output)
		awesome_error.CheckErr(err)
		return
	}
	workDir = string(match[1])
	return
}
