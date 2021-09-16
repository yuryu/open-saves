package blobref

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChecksums_CRC32C(t *testing.T) {
	var c Checksums
	c.SetCRC32C(uint32(0xfedcba98))
	assert.Equal(t, uint32(0xfedcba98), c.GetCRC32C())
}
