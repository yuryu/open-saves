package server

import (
	"bytes"
	"hash"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func updateChecksums(b []byte, md5 hash.Hash, crc32c hash.Hash32) (hash.Hash, hash.Hash32) {
	// Write never returns an error for Hash objects.
	// https://pkg.go.dev/hash#Hash
	_, _ = md5.Write(b)
	_, _ = crc32c.Write(b)
	return md5, crc32c
}

type withChecksums interface {
	GetHasCrc32C() bool
	GetCrc32C() uint32
	GetMd5Hash() []byte
}

func verifyChecksumsIfPresent(p withChecksums, md5 []byte, crc32c uint32) error {
	if len(p.GetMd5Hash()) != 0 {
		if !bytes.Equal(p.GetMd5Hash(), md5) {
			return status.Errorf(codes.DataLoss,
				"MD5 hash values didn't match: provided[%v], calculated[%v]",
				p.GetMd5Hash(), md5)
		}
	}
	if p.GetHasCrc32C() {
		if p.GetCrc32C() != crc32c {
			return status.Errorf(codes.DataLoss,
				"CRC32C checksums didn't match: provided[%v], calculated[%v]",
				p.GetCrc32C(), crc32c)
		}
	}
	return nil
}
