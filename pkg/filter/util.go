package filter

import "strings"

func ParseRule(rule string) (object Object) {
	var remained string
	var contains bool
	object.Package, object.FunctionName, contains = split(rule, ":::")
	if contains {
		return
	}
	remained, object.FunctionName, _ = split(rule, "::")
	object.Package, object.Filepath, _ = split(remained, ":")
	return
}

func split(rule, sep string) (first, second string, contains bool) {
	if strings.Contains(rule, sep) {
		splits := strings.Split(rule, sep)
		first = strings.TrimSpace(splits[0])
		second = splits[1]
		contains = true
	} else {
		first = rule
	}
	return
}
