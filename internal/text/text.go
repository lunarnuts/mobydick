package text

import (
	"bytes"
	"math/rand"
	"mobydick/internal/entity"
	"mobydick/internal/utils"
	"unicode"
)

// ProcessText читает содержиммое файла и возвращает список частот слов
func ProcessText(content []byte) []entity.WordFrequency {
	var wfs []entity.WordFrequency
	var buf bytes.Buffer

	for _, r := range content {
		// берем в буфер только буквы, имитируем команду tr -cs 'a-zA-Z' '[\n*]
		if isAlphabetical(rune(r)) {
			buf.WriteRune(unicode.ToLower(rune(r)))
		} else {
			// если встретили не букву, то проверяем буфер т.к. может быть пустой при встрече нескольких не букв подряд
			if buf.Len() > 0 {
				word := utils.BytesToRunes(buf.Bytes())
				//если есть слово в списке частот, то увеличиваем счетчик, иначе добавляем в список
				if index := wordIndex(word, wfs); index >= 0 {
					wfs[index].Count++
				} else {
					wfs = append(wfs, entity.WordFrequency{Word: word, Count: 1})
				}
				buf.Reset()
			}
		}
	}
	// может быть остаток в буфере
	if buf.Len() > 0 {
		word := utils.BytesToRunes(buf.Bytes())
		index := wordIndex(word, wfs)
		if index >= 0 {
			wfs[index].Count++
		} else {
			wfs = append(wfs, entity.WordFrequency{Word: word, Count: 1})
		}
	}

	return wfs
}

func isAlphabetical(r rune) bool {
	return r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z'
}

// wordIndex возвращает индекс слова в списке частот слов, если слово не найдено, то возвращает -1
func wordIndex(word []rune, wordFrequencies []entity.WordFrequency) int {
	for i, wf := range wordFrequencies {
		if utils.RunesEqual(word, wf.Word) {
			return i
		}
	}
	return -1
}

// FindTopNWords возвращает слайс с наиболее частыми N словами
func FindTopNWords(wfs []entity.WordFrequency, n int) []entity.WordFrequency {
	quickSort(wfs)
	if len(wfs) > n {
		return wfs[:n]
	}
	return wfs
}

// имплементация quicksort по убыванию
func quickSort(a []entity.WordFrequency) {
	if len(a) < 2 {
		return
	}

	l, r := 0, len(a)-1
	p := rand.Int() % len(a)
	a[p], a[r] = a[r], a[p]

	for i := range a {
		if a[i].Count > a[r].Count {
			a[i], a[l] = a[l], a[i]
			l++
		}
	}

	a[l], a[r] = a[r], a[l]
	quickSort(a[:l])
	quickSort(a[l+1:])
}
