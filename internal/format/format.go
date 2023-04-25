package format

import (
	"mobydick/internal/entity"
	"mobydick/internal/utils"
	"os"
	"sort"
)

// PrintWords форматирует вывод
func PrintWords(words []entity.WordFrequency) {
	sort.Slice(words, func(i, j int) bool {
		// имитирую sort -nr, если слова одинаковой частоты, то сортировать обратно алфавиту
		if words[i].Count == words[j].Count {
			a, b := words[i].Word, words[j].Word
			c, d := 0, 0
			for c < len(a) && d < len(b) {
				if a[c] != b[d] {
					return a[c] > b[d]
				}
				c++
				d++
			}
			return len(a) > len(b)
		}
		return words[i].Count > words[j].Count
	})

	for _, word := range words {
		s := utils.IntToRunes(word.Count)
		for len(s) < 7 {
			s = append([]rune{' '}, s...)
		}
		s = append(s, ' ')
		s = append(s, word.Word...)
		s = append(s, '\n')
		printRunes(s)
	}

}

func printRunes(r []rune) {
	for _, c := range r {
		_, err := os.Stdout.Write([]byte{byte(c)})
		if err != nil {
			return
		}
	}
}
