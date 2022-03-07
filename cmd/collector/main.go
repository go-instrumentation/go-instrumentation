package main

import (
	"encoding/json"
	go_instrumentation "github.com/go-instrumentation/go-instrumentation"
	"github.com/go-instrumentation/go-instrumentation/db"
	"github.com/go-instrumentation/go-instrumentation/db/model"
	"github.com/go-instrumentation/go-instrumentation/pkg/fingerprint"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/awesome_libs/log"
	"os"
)

func updatePackage(args []string, f flag) (err error) {
	compileFlag, err := go_instrumentation.ParseCompileFlag(args)
	if err != nil {
		return
	}
	if compileFlag.PkgPath == "" {
		return
	}
	db.MustPkgNotExistsInImportMap(compileFlag.PkgPath)
	compileCommand, err := json.Marshal(args)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	pkg := model.Package{
		PkgPath:        compileFlag.PkgPath,
		CfgPath:        compileFlag.ImportCfgPath,
		ImportPath:     compileFlag.PkgPath,
		BinaryPath:     compileFlag.Output,
		CompileCommand: compileCommand,
	}
	if f.FingerprintTool != "" {
		fp, err := fingerprint.Fingerprint(f.FingerprintTool, compileFlag.PkgPath, compileFlag.Output)
		if err != nil {
			return err
		}
		pkg.Fingerprint = fp
	}
	err = db.CreatePkg(pkg)
	if err != nil {
		return
	}
	return
}

func main() {
	log.Logger.Infof("origin args: %q", os.Args)
	args := os.Args[1:]
	args, flag := parseCliArgs(args)
	log.Logger.Debugf("%+v", flag)
	db.Init(flag.Dsn)
	cmd, args := go_instrumentation.ParseCommand(args)
	go_instrumentation.ForwardCommandPanicOnError(args)
	if cmd == "compile" {
		awesome_error.CheckFatal(updatePackage(args, flag))
	}
	go_instrumentation.Finish()
	os.Exit(0)
}
