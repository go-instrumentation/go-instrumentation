package fingerprint

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFingerprint(t *testing.T) {
	fingerprint, err := Fingerprint("/home/st0n3/go/bin/fingerprint", "context", "../../test/data/_pkg_.a")
	assert.NoError(t, err)
	assert.Equal(t, "c22b4c4d87217fe5", fingerprint)
}
