package utils

import (
	"fmt"
	"math/rand"
	"testing"
)

func Test_toSNAFU(t *testing.T) {
	for i := 0; i < 100; i++ {
		n := int(rand.Uint32())
		t.Run("", func(t *testing.T) {
			snafu := ToSNAFU(n)
			m := FromSNAFU(snafu)
			if m != n {
				msg := fmt.Sprintf("unexpected result in n = %d", n)
				t.Fatal(msg)
			}
		})
	}
}
