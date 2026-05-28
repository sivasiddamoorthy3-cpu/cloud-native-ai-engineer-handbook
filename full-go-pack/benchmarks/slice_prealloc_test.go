
package benchmarks

import "testing"

func BenchmarkSlicePrealloc(b *testing.B) {
    for i := 0; i < b.N; i++ {
        arr := make([]int, 0, 1000)

        for j := 0; j < 1000; j++ {
            arr = append(arr, j)
        }
    }
}
