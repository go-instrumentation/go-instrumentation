package filter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_regexMatch(t *testing.T) {
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
			name: "normal",
			args: args{
				rule: "main:.*\\.go::func",
				targetObject: Object{
					Package:      "main",
					Filepath:     "a.go",
					FunctionName: "func",
				},
			},
			wantMatch: true,
		},
		{
			name: "empty filepath",
			args: args{
				rule: "main:::func",
				targetObject: Object{
					Package:      "main",
					Filepath:     "",
					FunctionName: "func",
				},
			},
			wantMatch: true,
		},
		{
			name: "real world",
			args: args{
				rule: "github.com/containerd/containerd/.*",
				targetObject: Object{
					Package:      "vendor/github.com/containerd/containerd/a",
					Filepath:     "a",
					FunctionName: "func",
				},
			},
			wantMatch: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.wantMatch, regexMatch(tt.args.rule, tt.args.targetObject), "regexMatch(%v, %v)", tt.args.rule, tt.args.targetObject)
		})
	}
}
