// Popcount app exercises the popcount package
package popcount

import (
	"fmt"
	"testing"
)

func BenchmarkSelfConcatOperator1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fmt.Printf("popcount of %v (%b) is %d\n", n, n, PopCount(uint64(n)))
	}
}
