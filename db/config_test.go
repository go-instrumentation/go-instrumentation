package db

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPrebuilt(t *testing.T) {
	{
		prebuilt, err := IsPrebuilt()
		assert.NoError(t, err)
		assert.False(t, prebuilt)
	}
	assert.NoError(t, Prebuilt())
	{
		prebuilt, err := IsPrebuilt()
		assert.NoError(t, err)
		assert.True(t, prebuilt)
	}
}
