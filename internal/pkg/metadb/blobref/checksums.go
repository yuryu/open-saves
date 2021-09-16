package blobref

const (
	md5PropertyName    = "MD5"
	crc32cPropertyName = "CRC32C"
)

// Checksums is a struct for blob checksums.
type Checksums struct {
	// MD5 is the MD5 hash value of the associated blob object.
	MD5 []byte `datastore:",noindex"`
	// CRC32C is the CRC32C checksum of the associated blob object.
	// CRC32C uses the Castagnoli polynomial.
	// This needs to be int32 because Datastore doesn't support unsigned integers...
	// Use SetCRC32C() and GetCRC32C() for easier access.
	CRC32C int32 `datastore:",noindex"`
}

// SetCRC32C sets v to CRC32C.
func (c *Checksums) SetCRC32C(v uint32) {
	c.CRC32C = int32(v)
}

// GetCRC32C returns CRC32C.
func (c *Checksums) GetCRC32C() uint32 {
	return uint32(c.CRC32C)
}
