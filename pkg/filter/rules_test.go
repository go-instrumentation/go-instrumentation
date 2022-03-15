package filter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRule_Match(t *testing.T) {
	assert.False(t, RuleDenyGolang.Allow(Object{Package: "runtime"}))
	assert.False(t, RuleDenyPbDotGo.Allow(Object{Filepath: "a.pb.go"}))
	assert.True(t, RuleDenyPbDotGo.Allow(Object{Filepath: "main.go"}))
}
