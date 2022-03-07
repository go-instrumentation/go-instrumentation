package filter

import "strings"

func parseRule(rule string) (pkg, functionName string) {
	if strings.Contains(rule, ":") {
		splits := strings.Split(rule, ":")
		pkg = splits[0]
		functionName = splits[1]
	} else {
		pkg = rule
	}
	return
}
