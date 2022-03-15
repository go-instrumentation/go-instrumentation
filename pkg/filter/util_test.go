package filter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseRule(t *testing.T) {
	type args struct {
		rule string
	}
	tests := []struct {
		name       string
		args       args
		wantObject Object
	}{
		{
			name: "normal",
			args: args{
				rule: "main:*.go::func",
			},
			wantObject: Object{
				Package:      "main",
				Filepath:     "*.go",
				FunctionName: "func",
			},
		},
		{
			name: "empty package",
			args: args{
				rule: ":*.go::func",
			},
			wantObject: Object{
				Package:      "",
				Filepath:     "*.go",
				FunctionName: "func",
			},
		},
		{
			name: "empty filepath",
			args: args{
				rule: "main:::func",
			},
			wantObject: Object{
				Package:      "main",
				Filepath:     "",
				FunctionName: "func",
			},
		},
		{
			name: "empty functionName",
			args: args{
				rule: "main:*.go",
			},
			wantObject: Object{
				Package:      "main",
				Filepath:     "*.go",
				FunctionName: "",
			},
		},
		{
			name: "only package",
			args: args{
				rule: "main",
			},
			wantObject: Object{
				Package:      "main",
				Filepath:     "",
				FunctionName: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.wantObject, ParseRule(tt.args.rule), "ParseRule(%v)", tt.args.rule)
		})
	}
}
