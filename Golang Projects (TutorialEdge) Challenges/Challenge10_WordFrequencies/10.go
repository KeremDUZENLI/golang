package Challenge10_WordFrequencies

import (
	"sort"
	"strings"
)

type Word struct {
	Word  string
	Count int
}

func CountWords(text string) map[string]int {
	clean := strings.Replace(text, ",", "", -1)
	clean = strings.Replace(clean, ".", "", -1)

	splt := strings.Split(clean, " ")
	sort.Strings(splt)

	dict := make(map[string]int)

	for _, word := range splt {
		_, exist := dict[word]

		if exist {
			dict[word]++
		} else {
			dict[word] = 1
		}
	}

	return dict
}

func Top5Words(wordmap map[string]int) []Word {
	var liste []Word

	for word, count := range wordmap {
		liste = append(liste, Word{Word: word, Count: count})
	}

	sort.Slice(liste, func(i int, j int) bool {
		return liste[i].Count > liste[j].Count
	})
	return liste[:5]
}

func Test() []Word {
	text := "A A A A A B B B B C C C D D E"
	return Top5Words(CountWords(text))
}
