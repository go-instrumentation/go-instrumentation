package go_instrumentation

import (
	"fmt"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/awesome_libs/log"
	"golang.org/x/xerrors"
	"os"
	"os/exec"
	"strings"
)

func ParseCommand(args []string) (cmd string, newArgs []string) {
	newArgs = args
	binary := args[0]
	cmd = binary[strings.LastIndex(binary, "/")+1:]
	return
}

func ForwardCommand(args []string) error {
	path := args[0]
	args = args[1:]
	cmd := exec.Command(path, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	quotedArgs := fmt.Sprintf("%+q", args)
	log.Logger.Debugf("forwarding command `%s %s`", path, quotedArgs[1:len(quotedArgs)-1])
	return cmd.Run()
}

func ForwardCommandPanicOnError(args []string) {
	err := ForwardCommand(args)
	var exitErr *exec.ExitError
	if err != nil {
		awesome_error.CheckErr(err)
		if xerrors.As(err, &exitErr) {
			os.Exit(exitErr.ExitCode())
		} else {
			log.Logger.Fatal(err)
		}
	}
}
