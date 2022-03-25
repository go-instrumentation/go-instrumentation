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
	assert.True(t, Regex{AllowList: []string{"main"}}.Allow(Object{Package: "main", Filepath: "a.go"}))
	assert.False(t, Regex{AllowList: []string{"main", "pkg:a.go"}}.Allow(Object{Filepath: "b.go"}))
	assert.False(t, Regex{DenyList: []string{"k8s.io/cri-api/pkg/apis/runtime/v1:::&.*\\.MarshalToSizedBuffer"}}.Allow(
		Object{
			Package:      "k8s.io/cri-api/pkg/apis/runtime/v1",
			Filepath:     "a.go",
			FunctionName: "&ContainerMetadata.MarshalToSizedBuffer",
		},
	))
	assert.True(t, Regex{AllowList: []string{"github.com/containerd/containerd/pkg/cri/server:container_create.go::.*CreateContainer.*"}}.Allow(
		Object{
			Package:      "github.com/containerd/containerd/pkg/cri/server",
			Filepath:     "container_create.go",
			FunctionName: "&criService.CreateContainer",
		},
	))
}
