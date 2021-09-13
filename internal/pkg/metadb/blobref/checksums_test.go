package blobref

import (
	"crypto/md5"
	"hash/crc32"
	"testing"

	"cloud.google.com/go/datastore"
	"github.com/stretchr/testify/assert"
)

func TestChecksums_SaveLoad(t *testing.T) {
	const testData = "This is a test."
	md5 := md5.New()
	md5.Write([]byte(testData))
	crc32c := crc32.New(crc32.MakeTable(crc32.Castagnoli))
	crc32c.Write([]byte(testData))

	checksums := &Checksums{
		MD5Hash: md5.Sum(nil),
		CRC32C:  crc32c.Sum32(),
	}
	ps, err := checksums.Save()
	assert.NoError(t, err, "Save shuold not return err.")
	expected := datastore.PropertyList{
		datastore.Property{
			Name:    "MD5Hash",
			Value:   md5.Sum(nil),
			NoIndex: true,
		},
		datastore.Property{
			Name:    "CRC32C",
			Value:   int64(crc32c.Sum32()),
			NoIndex: true,
		},
	}
	if assert.NotEmpty(t, ps) {
		assert.ElementsMatch(t, expected, ps)
	}

	loaded := new(Checksums)
	if assert.NoError(t, loaded.Load(expected)) {
		assert.Equal(t, md5.Sum(nil), loaded.MD5Hash)
		assert.Equal(t, crc32c.Sum32(), loaded.CRC32C)
	}
}
