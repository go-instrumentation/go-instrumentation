package filter

import "strings"

type Prefix struct {
	AllowList []string
	DenyList  []string
}

var PrefixFilter Prefix

func (f Prefix) Match(pkg, functionName string) (result bool) {
	for _, deny := range f.DenyList {
		if prefixMatch(deny, pkg, functionName) {
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
		if prefixMatch(rule, pkg, functionName) {
			return true
		}
	}
	return false
}

func prefixMatch(rule, pkg, functionName string) (match bool) {
	rulePkg, ruleFunctionName := parseRule(rule)
	if strings.HasSuffix(rulePkg, "*") { // means no functionName
		if strings.HasPrefix(pkg, strings.TrimSuffix(rulePkg, "*")) {
			return true
		}
	} else { // may have functionName
		if rulePkg == pkg {
			if ruleFunctionName == "" {
				return true
			} else {
				return ruleFunctionName == functionName
			}
		}
	}
	return false
}
