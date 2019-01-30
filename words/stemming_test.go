package words_test

import (
	"testing"

	"github.com/blainsmith/goreds/words"
)

func TestStems(t *testing.T) {
	testwords := []string{"Trying", "counting", "ten", "count", "thirteen", "filing"}

	stems := words.StemsArray(testwords)
	if len(stems) != 6 {
		t.Error("expected 7 word keys")
	}
}

func BenchmarkStems(b *testing.B) {
	testwords := []string{"Trying", "counting", "ten", "count", "thirteen", "filing"}

	for i := 0; i < b.N; i++ {
		words.StemsArray(testwords)
	}
}
