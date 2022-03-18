package filter

import "strings"

type Prefix struct {
	Base
	AllowList []string
	DenyList  []string
}

func (f Prefix) Allow(targetObject Object) (allow bool) {
	for _, rule := range f.DenyList {
		if prefixMatch(rule, targetObject) {
			debug(f, rule, targetObject, allow)
			return
		}
	}
	if f.AllowList == nil {
		allow = true
		debug(f, "AllowList == nil", targetObject, allow)
		return
	}
	for _, rule := range f.AllowList {
		if rule == "*" {
			allow = true
			debug(f, rule, targetObject, allow)
			return
		}
		if prefixMatch(rule, targetObject) {
			allow = true
			debug(f, rule, targetObject, allow)
			return
		}
	}
	debug(f, "no matching rule", targetObject, allow)
	return
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
