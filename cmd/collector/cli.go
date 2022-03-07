package main

import (
	"fmt"
	"github.com/ssst0n3/awesome_libs/log"
	"os"
)

type flag struct {
	Dsn             string
	FingerprintTool string
}

const (
	errorMsgFlagProvideButNotDefined = "flag provided but not defined: %s"
)

func parseCliArgs(args []string) (newArgs []string, f flag) {
	newArgs = f.parseArgs(args)
	log.Logger.Debugf("newArgs: %+v", newArgs)
	if f.Dsn == "" {
		f.Dsn = "/tmp/go_instrumentation.sqlite"
	}
	return
}

func (f *flag) parseArgs(args []string) (newArgs []string) {
	var pass bool
	for i, arg := range args {
		if pass {
			pass = false
			continue
		}
		switch arg {
		case "-h":
			usage("")
		case "-dsn":
			pass = true
			if len(args) < i+2 {
				usage(fmt.Sprintf(errorMsgFlagProvideButNotDefined, arg))
			}
			f.Dsn = args[i+1]
		case "-fp":
			pass = true
			if len(args) < i+2 {
				usage(fmt.Sprintf(errorMsgFlagProvideButNotDefined, arg))
			}
			f.FingerprintTool = args[i+1]
		default:
			newArgs = append(newArgs, arg)
		}
	}
	return
}

func usage(errorMsg string) {
	if len(errorMsg) > 0 {
		fmt.Printf("Incorrect Usage: %s\n", errorMsg)
	}
	fmt.Println("Usage: ")
	os.Exit(1)
}
