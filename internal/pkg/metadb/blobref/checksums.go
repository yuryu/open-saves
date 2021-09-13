package blobref

import (
	"cloud.google.com/go/datastore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const crc32cPropertyName = "CRC32C"

// Checksums is a struct for blob checksums.
type Checksums struct {
	// MD5Hash is the MD5 hash value of the associated blob object.
	MD5Hash []byte `datastore:",noindex"`
	// CRC32C is the CRC32C checksum of the associated blob object.
	// CRC32C uses the Castagnoli polynomial.
	CRC32C uint32 `datastore:"-"`
}

// Assert Checksums implements PropertyLoadSaver.
var _ datastore.PropertyLoadSaver = new(Checksums)

func (c *Checksums) Save() ([]datastore.Property, error) {
	properties, err := datastore.SaveStruct(c)
	if err != nil {
		return nil, err
	}
	properties = append(properties,
		datastore.Property{
			Name:    crc32cPropertyName,
			Value:   int64(c.CRC32C),
			NoIndex: true,
		},
	)
	return properties, nil
}

func (c *Checksums) Load(ps []datastore.Property) error {
	for i, p := range ps {
		if p.Name == crc32cPropertyName {
			if v, ok := p.Value.(int64); ok {
				ps[i] = ps[len(ps)-1]
				ps = ps[:len(ps)-1]
				c.CRC32C = uint32(v)
			} else {
				return status.Errorf(codes.Internal, "CRC32C property is not integer: %v", p.Value)
			}
		}
	}
	return datastore.LoadStruct(c, ps)
}
