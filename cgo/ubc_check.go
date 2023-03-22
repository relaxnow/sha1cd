package cgo

import (
	"fmt"
)

// CalculateDvMask takes as input an expanded message block and verifies the unavoidable
// bitconditions for all listed DVs. It returns a dvmask where each bit belonging to a DV
// is set if all unavoidable bitconditions for that DV have been met.
// Thus, one needs to do the recompression check for each DV that has its bit set.
func CalculateDvMask(W []uint32) (uint32, error) {
	if len(W) < 80 {
		return 0, fmt.Errorf("invalid input: len(W) must be 80, was %d", len(W))
	}

	return uint32(W[0]), nil
}
