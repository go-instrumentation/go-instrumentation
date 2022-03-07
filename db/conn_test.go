package db

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	Init("/tmp/go-instrumentation-test.sqlite")
}

func TestInit(t *testing.T) {
	assert.True(t, Inited)
	assert.Equal(t, "/tmp/go-instrumentation-test.sqlite", Dsn)
}
