package filter

import (
	"github.com/sirupsen/logrus"
	"github.com/ssst0n3/awesome_libs/log"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRule_Match(t *testing.T) {
	log.Logger.Level = logrus.DebugLevel
	assert.False(t, RuleDenyGolang.Allow(Object{Package: "runtime"}))
	assert.False(t, RuleDenyPbDotGo.Allow(Object{Filepath: "a.pb.go"}))
	assert.True(t, RuleDenyPbDotGo.Allow(Object{Filepath: "main.go"}))
	assert.True(t, RuleDenyTooManyDetails.Allow(Object{Package: "main"}))
	assert.False(t, RuleDenyTooManyDetails.Allow(Object{Package: "context"}))
}
