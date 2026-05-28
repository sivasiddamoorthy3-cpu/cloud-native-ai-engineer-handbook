
package benchmarks

import (
    "strings"
    "testing"
)

func BenchmarkStringBuilder(b *testing.B) {
    for i := 0; i < b.N; i++ {
        var sb strings.Builder

        for j := 0; j < 100; j++ {
            sb.WriteString("hello")
        }
    }
}
