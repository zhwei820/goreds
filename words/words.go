package words // import "github.com/blainsmith/goreds/words"

import "strings"

var stopwords = []string{
	"about", "above", "after", "again", "all", "also", "am", "an", "and", "another", "any", "are", "as", "at",
	"be", "because", "been", "before", "being", "below", "between", "both", "but", "by",
	"came", "can", "cannot", "come", "could",
	"did", "do", "does", "doing", "during",
	"each",
	"few", "for", "from", "further",
	"get", "got",
	"has", "had", "he", "have", "her", "here", "him", "himself", "his", "how",
	"if", "in", "into", "is", "it", "its", "itself",
	"like",
	"make", "many", "me", "might", "more", "most", "much", "must", "my", "myself",
	"never", "now",
	"of", "on", "only", "or", "other", "our", "ours", "ourselves", "out", "over", "own",
	"said", "same", "see", "should", "since", "so", "some", "still", "such",
	"take", "than", "that", "the", "their", "theirs", "them", "themselves", "then", "there", "these", "they", "this", "those", "through", "to", "too",
	"under", "until", "up",
	"very",
	"was", "way", "we", "well", "were", "what", "where", "when", "which", "while", "who", "whom", "with", "would", "why",
	"you", "your", "yours", "yourself",
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"$", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0", "_",
	"",
}

var chstopwords = []string{
	",", ".", "?", "!", "\"", " ", "@", "，", "。", "、", "？", "！", "：", "“", "”", "；", "　", "（", "）", "《", "》", "~", "*", "<", ">",
	"/", "\\", "|", "-", "_", "+", "=", "&", "^", "%", "#", "`", ";", "$", "￥", "‘", "’", "〉", "〈", "…", "＞", "＜", "＠", "＃", "＄", "％",
	"︿", "＆", "＊", "＋", "～", "｜", "［", "］", "｛", "｝",
	"",
}

func Split(text string) []string {
	return strings.Split(text, " ")
}

func Count(words []string) map[string]int {
	countMap := make(map[string]int)

	for _, w := range words {
		countMap[w] = countMap[w] + 1
	}

	return countMap
}

func StripStopWords(words []string) []string {
	for _, sw := range stopwords {
		var ii = 0
		for i, w := range words {
			if w == sw {
				words = append(words[:i-ii], words[i+1-ii:]...)
				ii += 1
			}
		}
	}

	return words
}

func StripChStopWords(words []string) []string {
	for _, sw := range chstopwords {
		var ii = 0
		for i, w := range words {
			if w == sw {
				words = append(words[:i-ii], words[i+1-ii:]...)
				ii += 1
			}
		}
	}

	return words
}
