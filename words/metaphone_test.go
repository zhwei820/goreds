package words_test

import (
	"testing"

	"blainsmith.com/go/goreds/words"
)

func TestMetaphoneMap(t *testing.T) {
	testwords := []string{"Trying", "count", "ten", "count", "thirteen"}

	metaphones := words.MetaphoneMap(testwords)
	if len(metaphones) != 4 {
		t.Error("expected 7 word keys")
	}
}

func BenchmarkMetaphoneMap(b *testing.B) {
	testwords := []string{"Trying", "count", "ten", "count", "thirteen"}

	for i := 0; i < b.N; i++ {
		words.MetaphoneMap(testwords)
	}
}

func TestMetaphoneArray(t *testing.T) {
	testwords := []string{"Trying", "count", "ten", "count", "thirteen"}

	metaphones := words.MetaphoneArray(testwords)
	if len(metaphones) != 4 {
		t.Error("expected 7 word keys")
	}
}

func BenchmarkMetaphoneArray(b *testing.B) {
	testwords := []string{"Trying", "count", "ten", "count", "thirteen"}

	for i := 0; i < b.N; i++ {
		words.MetaphoneArray(testwords)
	}
}

func TestMetaphoneKeys(t *testing.T) {
	testwords := []string{"Trying", "count", "ten", "count", "thirteen"}

	keys := words.MetaphoneKeys("namespace", testwords)
	if len(keys) != 4 {
		t.Error("expected 7 word keys")
	}
}

func BenchmarkMetaphoneKeys(b *testing.B) {
	testwords := []string{"Trying", "count", "ten", "count", "thirteen"}

	for i := 0; i < b.N; i++ {
		words.MetaphoneKeys("namespace", testwords)
	}
}
