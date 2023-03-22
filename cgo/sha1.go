package cgo

import (
	"crypto"
	"hash"
)

const (
	Size      = 20
	BlockSize = 64
)

func init() {
	crypto.RegisterHash(crypto.SHA1, New)
}

func New() hash.Hash {
	d := new(digest)
	d.Reset()
	return d
}

type SHA1_CTX struct {
}

type digest struct {
	ctx SHA1_CTX
}

func (d *digest) sum() ([]byte, bool) {
	b := make([]byte, Size)
	return b, false
}

func (d *digest) Sum(in []byte) []byte {
	d0 := *d // use a copy of d to avoid race conditions.
	h, _ := d0.sum()
	return append(in, h...)
}

func (d *digest) CollisionResistantSum(in []byte) ([]byte, bool) {
	d0 := *d // use a copy of d to avoid race conditions.
	h, c := d0.sum()
	return append(in, h...), c
}

func (d *digest) Reset() {
}

func (d *digest) Size() int { return Size }

func (d *digest) BlockSize() int { return BlockSize }

func Sum(data []byte) ([]byte, bool) {
	d := New().(*digest)
	d.Write(data)

	return d.sum()
}
