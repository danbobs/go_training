// Popcount app exercises the popcount package
package popcount

import (
	"testing"
)

func BenchmarkPopCount(b *testing.B) {
	for n := 0; n < b.N; n++ {
		PopCount3(uint64(n))
	}
}
