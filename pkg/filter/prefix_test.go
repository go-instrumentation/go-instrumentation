package filter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_prefixMatch(t *testing.T) {
	type args struct {
		rule         string
		targetObject Object
	}
	tests := []struct {
		name      string
		args      args
		wantMatch bool
	}{
		{
			name: "match",
			args: args{
				rule: "pkg:&a.b",
				targetObject: Object{
					Package:      "pkg",
					FunctionName: "&a.b",
				},
			},
			wantMatch: true,
		},
		{
			name: "no :",
			args: args{
				rule: "pkg",
				targetObject: Object{
					Package:      "pkg",
					FunctionName: "&a.b",
				},
			},
			wantMatch: true,
		},
		{
			name: "*",
			args: args{
				rule: "pkg*",
				targetObject: Object{
					Package:      "pkg.pkg",
					FunctionName: "&a.b",
				},
			},
			wantMatch: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.wantMatch, prefixMatch(tt.args.rule, tt.args.targetObject), "match(%v, %v, %v)", tt.args.rule, tt.args.targetObject)
		})
	}
}
