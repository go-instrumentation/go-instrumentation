package fingerprint

import (
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"os/exec"
	"strings"
)

func Fingerprint(execute, pkgPath, binaryPath string) (fingerprint string, err error) {
	cmd := exec.Command(execute, "-pkg", pkgPath, "-file", binaryPath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	fingerprint = strings.TrimSpace(string(output))
	return
}
