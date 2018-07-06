package words // import "blainsmith.com/go/goreds/words"

import "github.com/a2800276/porter"

func StemsArray(words []string) []string {
	for i, w := range words {
		words[i] = porter.Stem(w)
	}

	return words
}
