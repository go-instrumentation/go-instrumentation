package filter

type Filter interface {
	Match(pkg, functionName string) (result bool)
}
