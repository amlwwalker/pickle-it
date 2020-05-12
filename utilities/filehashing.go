package utilities

import (
	"github.com/kalafut/imohash"
)

// UniqueFileHash creats a fast hash of a file. It's not bullet proof (could cause a collision, but in practice unlikely) but its fast
func UniqueFileHash(src string) ([16]byte, error) {
	hash, err := imohash.SumFile(src)
	if err != nil {
		return [16]byte{}, err
	}
	return hash, nil
}
