package ngrams

import (
	"fmt"
	"strings"
)

var ErrParam = fmt.Errorf("invalid value of")

// MakeRange generates ngrams with len=min..max
func MakeRange(word string, min, max int) ([]string, error) {
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

	runes := []rune(word)
	wordLen := len(runes)
	maxN := minOf(wordLen, max)
	result := make([]string, 0, ngramCnt(len(runes), min, maxN))
	for n := min; n <= maxN; n++ {
		items, err := FromRunes(runes, n)
		if err != nil {
			return nil, err
		}
		result = append(result, items...)
	}

	return result, nil
}

// From generates ngrams from a string with len=n
func From(word string, n int) ([]string, error) {
	return FromRunes([]rune(word), n)
}

// FromRunes generates ngrams from a rune slice with len=n
func FromRunes(runes []rune, n int) ([]string, error) {
	if n < 1 {
		return nil, fmt.Errorf("%w: n cannot be < 1", ErrParam)
	}

	if n > len(runes) {
		return nil, nil
	}

	if n == len(runes) {
		return []string{string(runes)}, nil
	}

	cnt := ngramCnt(len(runes), n, n)
	builders := make([]strings.Builder, cnt)

	for i, r := range runes {
		for j := i - n + 1; j <= i; j++ {
			if j < 0 {
				continue
			}
			if j >= cnt {
				break
			}

			builders[j].WriteRune(r)
		}
	}

	result := make([]string, cnt)
	for i, b := range builders {
		result[i] = b.String()
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
