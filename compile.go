package go_instrumentation

// CmdFlags
// https://github.com/golang/go/blob/master/src/cmd/compile/internal/base/flag.go
type CmdFlags struct {
	PkgPath       string `names:"-p"`
	Output        string `names:"-o"`
	ImportCfgPath string `names:"-importcfg"`
}

func ParseCompileFlag(args []string) (f CmdFlags, err error) {
	for i, arg := range args {
		switch arg {
		case "-p":
			f.PkgPath = args[i+1]
		case "-o":
			f.Output = args[i+1]
		case "-importcfg":
			f.ImportCfgPath = args[i+1]
		}
	}
	return
}
