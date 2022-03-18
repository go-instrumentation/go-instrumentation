package filter

import "regexp"

type Regex struct {
	Base
	AllowList []string
	DenyList  []string
}

func (f Regex) Allow(targetObject Object) (allow bool) {
	for _, rule := range f.DenyList {
		if regexMatch(rule, targetObject) {
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
		if regexMatch(rule, targetObject) {
			allow = true
			debug(f, rule, targetObject, allow)
		}
	}
	return
}

func regexMatch(rule string, targetObject Object) (match bool) {
	ruleObject := ParseRule(rule)
	return regexp.MustCompile(ruleObject.Package).MatchString(targetObject.Package) &&
		regexp.MustCompile(ruleObject.Filepath).MatchString(targetObject.Filepath) &&
		regexp.MustCompile(ruleObject.FunctionName).MatchString(targetObject.FunctionName)
}
