package filter

type Filter interface {
	Allow(pkg, functionName string) (result bool)
}
