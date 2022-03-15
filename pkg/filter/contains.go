package filter

import "strings"

type Contains struct {
	AllowList []string
	DenyList  []string
}

func (f Contains) Allow(pkg, functionName string) (result bool) {
	for _, deny := range f.DenyList {
		if contains(deny, pkg, functionName) {
			return false
		}
	}
	if f.AllowList == nil {
		return true
	}
	for _, rule := range f.AllowList {
		if contains(rule, pkg, functionName) {
			return true
		}
	}
	return false
}

func contains(rule, pkg, functionName string) (match bool) {
	rulePkg, ruleFunctionName := parseRule(rule)
	return strings.Contains(pkg, rulePkg) && strings.Contains(functionName, ruleFunctionName)
}
