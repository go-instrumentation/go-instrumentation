package filter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRule_Match(t *testing.T) {
	assert.False(t, RuleDenyGolang.Allow("runtime", ""))
}
