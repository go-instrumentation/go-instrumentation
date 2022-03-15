package filter

import "strings"

type Contains struct {
	AllowList []string
	DenyList  []string
}

func (f Contains) Allow(targetObject Object) (allow bool) {
	for _, deny := range f.DenyList {
		if contains(deny, targetObject) {
			return false
		}
	}
	if f.AllowList == nil {
		return true
	}
	for _, rule := range f.AllowList {
		if contains(rule, targetObject) {
			return true
		}
	}
	return false
}

func contains(ruleRaw string, targetObject Object) (match bool) {
	ruleObject := ParseRule(ruleRaw)
	return strings.Contains(targetObject.Package, ruleObject.Package) && strings.Contains(targetObject.FunctionName, ruleObject.FunctionName)
}
