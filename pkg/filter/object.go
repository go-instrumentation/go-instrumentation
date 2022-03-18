package filter

import "fmt"

type Object struct {
	Package      string
	Filepath     string
	FunctionName string
}

func (o Object) String() string {
	return fmt.Sprintf("%s:%s::%s", o.Package, o.Filepath, o.FunctionName)
}
