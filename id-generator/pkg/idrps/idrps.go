package idrps

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"sync/atomic"
	"time"
)

type ID []byte

func (i ID) Hex() string {
	return hex.EncodeToString(i)
}

func (i ID) Base64() string {
	return base64.RawStdEncoding.EncodeToString(i)
}

type IDRPS struct {
	randomSize int
	counter    uint32
}

// NewIDRPS is a function to build NewIDRPS
func NewIDRPS(randomSize int) (*IDRPS, error) {

	counter := make([]byte, 4)

	_, err := rand.Read(counter)

	if err != nil {
		return nil, err
	}

	return &IDRPS{
		randomSize: randomSize,
		counter:    binary.LittleEndian.Uint32(counter),
	}, nil
}

// Generate as a function to create new IDs
func (i *IDRPS) Generate() (ID, error) {
	id := make(ID, i.randomSize+12)

	_, err := rand.Read(id[:i.randomSize])

	if err != nil {
		return nil, err
	}

	now := uint64(time.Now().Unix())
	binary.LittleEndian.PutUint64(id[i.randomSize:i.randomSize+8], now)

	counter := atomic.AddUint32(&(i.counter), 1)
	binary.LittleEndian.PutUint32(id[i.randomSize+8:i.randomSize+12], counter)

	return id, nil
}
