package filter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_prefixMatch(t *testing.T) {
	type args struct {
		rule         string
		pkg          string
		functionName string
	}
	tests := []struct {
		name      string
		args      args
		wantMatch bool
	}{
		{
			name: "match",
			args: args{
				rule:         "pkg:&a.b",
				pkg:          "pkg",
				functionName: "&a.b",
			},
			wantMatch: true,
		},
		{
			name: "no :",
			args: args{
				rule:         "pkg",
				pkg:          "pkg",
				functionName: "&a.b",
			},
			wantMatch: true,
		},
		{
			name: "*",
			args: args{
				rule:         "pkg*",
				pkg:          "pkg.pkg",
				functionName: "&a.b",
			},
			wantMatch: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.wantMatch, prefixMatch(tt.args.rule, tt.args.pkg, tt.args.functionName), "match(%v, %v, %v)", tt.args.rule, tt.args.pkg, tt.args.functionName)
		})
	}
}
