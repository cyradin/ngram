package ngrams

import (
	"fmt"
)

var ErrParam = fmt.Errorf("invalid value of")

// MakeRange generates ngrams with len=min..max
func MakeRange(word string, min, max int) ([][]rune, error) {
	if min < 1 {
		return nil, fmt.Errorf("%w min: cannot be < 1", ErrParam)
	}
	if max < 1 {
		return nil, fmt.Errorf("%w min: cannot be < 1", ErrParam)
	}
	if min > max {
		return nil, fmt.Errorf("%w min: cannot be > max", ErrParam)
	}

	if word == "" {
		return nil, nil
	}

	wordLen := len([]rune(word))
	maxN := minOf(wordLen, max)
	result := make([][]rune, 0, ngramCnt(wordLen, min, maxN))
	for n := min; n <= maxN; n++ {
		items, err := From(word, n)
		if err != nil {
			return nil, err
		}
		result = append(result, items...)
	}

	return result, nil
}

// From generates ngrams from a string with len=n
func From(word string, n int) ([][]rune, error) {
	return FromRunes([]rune(word), n)
}

// FromRunes generates ngrams from a rune slice with len=n
func FromRunes(runes []rune, n int) ([][]rune, error) {
	if n < 1 {
		return nil, fmt.Errorf("%w: n cannot be < 1", ErrParam)
	}

	if n > len(runes) {
		return nil, nil
	}

	if n == len(runes) {
		return [][]rune{runes}, nil
	}

	cnt := ngramCnt(len(runes), n, n)
	buf := make([]rune, n*cnt)

	for i, r := range runes {
		for j := i - n + 1; j <= i; j++ {
			if j < 0 {
				continue
			}
			if j >= cnt {
				break
			}

			index := j*n + (i - j)

			buf[index] = r
		}
	}

	result := make([][]rune, cnt)
	for i := 0; i < cnt; i++ {
		result[i] = buf[i*n : (i+1)*n]
	}

	return result, nil
}

func ngramCnt(l, min, max int) int {
	if min == max {
		return l - min + 1
	}

	var cnt int
	for i := min; i <= max; i++ {
		cnt += l - i + 1
	}
	return cnt
}

func minOf(a, b int) int {
	if a > b {
		return b
	}
	return a
}
