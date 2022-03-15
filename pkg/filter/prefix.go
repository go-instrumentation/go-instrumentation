package filter

import "strings"

type Prefix struct {
	AllowList []string
	DenyList  []string
}

var PrefixFilter Prefix

func (f Prefix) Allow(targetObject Object) (allow bool) {
	for _, rule := range f.DenyList {
		if prefixMatch(rule, targetObject) {
			return false
		}
	}
	if f.AllowList == nil {
		return true
	}
	for _, rule := range f.AllowList {
		if rule == "*" {
			return true
		}
		if prefixMatch(rule, targetObject) {
			return true
		}
	}
	return false
}

func prefixMatch(rule string, targetObject Object) (match bool) {
	ruleObject := ParseRule(rule)
	if strings.HasSuffix(ruleObject.Package, "*") { // means no functionName
		if strings.HasPrefix(targetObject.Package, strings.TrimSuffix(ruleObject.Package, "*")) {
			return true
		}
	} else { // may have functionName
		if ruleObject.Package == targetObject.Package {
			if ruleObject.FunctionName == "" {
				return true
			} else {
				return ruleObject.FunctionName == targetObject.FunctionName
			}
		}
	}
	return false
}
