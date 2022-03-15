package filter

import "regexp"

type Regex struct {
	AllowList []string
	DenyList  []string
}

func (f Regex) Allow(targetObject Object) (allow bool) {
	for _, rule := range f.DenyList {
		if regexMatch(rule, targetObject) {
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
		if regexMatch(rule, targetObject) {
			return true
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
