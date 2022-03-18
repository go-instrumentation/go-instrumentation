package filter

import (
	"strings"
)

type Contains struct {
	Base
	AllowList []string
	DenyList  []string
}

func (f Contains) Allow(targetObject Object) (allow bool) {
	for _, deny := range f.DenyList {
		if contains(deny, targetObject) {
			debug(f, deny, targetObject, allow)
			return
		}
	}
	if f.AllowList == nil {
		return true
	}
	for _, rule := range f.AllowList {
		if contains(rule, targetObject) {
			allow = true
			debug(f, rule, targetObject, allow)
			return
		}
	}
	debug(f, "no matching rule", targetObject, allow)
	return
}

func contains(ruleRaw string, targetObject Object) (match bool) {
	ruleObject := ParseRule(ruleRaw)
	return strings.Contains(targetObject.Package, ruleObject.Package) && strings.Contains(targetObject.FunctionName, ruleObject.FunctionName)
}
