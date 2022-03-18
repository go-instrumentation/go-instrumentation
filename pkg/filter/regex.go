package filter

import "regexp"

type Regex struct {
	Base
	AllowList []string
	DenyList  []string
}

func (f Regex) Allow(targetObject Object) (allow bool) {
	for _, rule := range f.DenyList {
		if regexMatch(rule, targetObject, false) {
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
		if regexMatch(rule, targetObject, true) {
			allow = true
			debug(f, rule, targetObject, allow)
		}
	}
	return
}

func regexMatch(rule string, targetObject Object, matchEmpty bool) (match bool) {
	ruleObject := ParseRule(rule)
	var matchPackage, matchFilepath, matchFunctionName bool
	if matchEmpty {
		matchFilepath = targetObject.Filepath == ""
		matchFunctionName = targetObject.FunctionName == ""
	}
	matchPackage = regexp.MustCompile(ruleObject.Package).MatchString(targetObject.Package)
	matchFilepath = matchFilepath || regexp.MustCompile(ruleObject.Filepath).MatchString(targetObject.Filepath)
	matchFunctionName = matchFunctionName || regexp.MustCompile(ruleObject.FunctionName).MatchString(targetObject.FunctionName)
	return matchPackage && matchFilepath && matchFunctionName
}
