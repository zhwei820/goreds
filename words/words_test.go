package words_test

import (
	"testing"

	"github.com/blainsmith/goreds/words"
)

func TestSplit(t *testing.T) {
	testwords := "this is a set of words to split apart"

	split := words.Split(testwords)
	if len(split) != 9 {
		t.Error("expected 9 words split")
	}
}

func BenchmarkSplit(b *testing.B) {
	testwords := "this is a set of words to split apart"

	for i := 0; i < b.N; i++ {
		words.Split(testwords)
	}
}

func TestCount(t *testing.T) {
	testwords := []string{"Trying", "to", "count", "to", "ten", "and", "then", "count", "to", "thirteen"}

	counts := words.Count(testwords)
	if len(counts) != 7 {
		t.Error("expected 7 word keys")
	}
}

func BenchmarkCount(b *testing.B) {
	testwords := []string{"Trying", "to", "count", "to", "ten", "and", "then", "count", "to", "thirteen"}

	for i := 0; i < b.N; i++ {
		words.Count(testwords)
	}
}

func TestStripStopWords(t *testing.T) {
	testwords := []string{"Trying", "to", "count", "to", "ten", "and", "then", "count", "to", "thirteen"}

	stripped := words.StripStopWords(testwords)
	if len(stripped) != 5 {
		t.Error("expected 5 word keys")
	}
}

func BenchmarkStripStopWords(b *testing.B) {
	testwords := []string{"Trying", "to", "count", "to", "ten", "and", "then", "count", "to", "thirteen"}

	for i := 0; i < b.N; i++ {
		words.StripStopWords(testwords)
	}
}
