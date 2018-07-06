package words // import "blainsmith.com/go/goreds/words"

import "github.com/dotcypress/phonetics"

func MetaphoneMap(words []string) map[string]string {
	metaphones := make(map[string]string)

	for _, w := range words {
		metaphones[w] = phonetics.EncodeMetaphone(w)
	}

	return metaphones
}

func MetaphoneArray(words []string) []string {
	metaphones := make(map[string]struct{})

	for _, w := range words {
		metaphones[phonetics.EncodeMetaphone(w)] = struct{}{}
	}

	words = words[0:0]
	for i, _ := range metaphones {
		words = append(words, i)
	}

	return words
}

func MetaphoneKeys(key string, words []string) []string {
	var keys []string

	for _, m := range MetaphoneArray(words) {
		keys = append(keys, key+":word:"+m)
	}

	return keys
}
