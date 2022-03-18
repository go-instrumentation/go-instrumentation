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
		if regexMatch(rule, targetObject) {
			allow = true
			debug(f, rule, targetObject, allow)
		}
	}
	return
}

func regexMatch(rule string, targetObject Object) (match bool) {
	ruleObject := ParseRule(rule)
	return (targetObject.Package == "" || regexp.MustCompile(ruleObject.Package).MatchString(targetObject.Package)) &&
		(targetObject.Filepath == "" || regexp.MustCompile(ruleObject.Filepath).MatchString(targetObject.Filepath)) &&
		(targetObject.FunctionName == "" || regexp.MustCompile(ruleObject.FunctionName).MatchString(targetObject.FunctionName))
}
