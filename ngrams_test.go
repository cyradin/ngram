package ngrams

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Benchmark_MakeRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MakeRange("orange", 1, 5)
	}
}

func Test_MakeRange(t *testing.T) {
	t.Run("must return error if min < 1", func(t *testing.T) {
		result, err := MakeRange("word", 0, 5)
		require.Error(t, err)
		require.Nil(t, result)
	})

	t.Run("must return error if max < 1", func(t *testing.T) {
		result, err := MakeRange("word", 5, 0)
		require.Error(t, err)
		require.Nil(t, result)
	})

	t.Run("must return error if min > max", func(t *testing.T) {
		result, err := MakeRange("word", 2, 1)
		require.Error(t, err)
		require.Nil(t, result)
	})

	t.Run("must return nil for an empty word", func(t *testing.T) {
		result, err := MakeRange("", 1, 5)
		require.NoError(t, err)
		require.Nil(t, result)
	})

	t.Run("must return valid ngrams", func(t *testing.T) {
		t.Run("min=3,max=3", func(t *testing.T) {
			result, err := MakeRange("orange", 3, 3)
			require.NoError(t, err)
			require.Equal(t, []string{"ora", "ran", "ang", "nge"}, result)
		})

		t.Run("min=1,max=5", func(t *testing.T) {
			result, err := MakeRange("orange", 1, 5)
			require.NoError(t, err)
			require.Equal(t, []string{
				"o", "r", "a", "n", "g", "e",
				"or", "ra", "an", "ng", "ge",
				"ora", "ran", "ang", "nge",
				"oran", "rang", "ange",
				"orang", "range",
			}, result)
		})
	})
}

func Benchmark_From_6_3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		From("qwerty", 3)
	}
}

func Benchmark_FromRunes_6_3(b *testing.B) {
	runes := []rune("qwerty")
	for i := 0; i < b.N; i++ {
		FromRunes(runes, 3)
	}
}

func Test_From(t *testing.T) {
	t.Run("must return error if n < 1", func(t *testing.T) {
		result, err := From("qwe", 0)
		require.Error(t, err)
		require.Nil(t, result)
	})

	t.Run("must return nil if n > word len", func(t *testing.T) {
		result, err := From("qwe", 5)
		require.NoError(t, err)
		require.Nil(t, result)
	})

	t.Run("must return word if n == word len", func(t *testing.T) {
		result, err := From("qwe", 3)
		require.NoError(t, err)
		require.Equal(t, []string{"qwe"}, result)
	})

	t.Run("must return correct set of ngrams", func(t *testing.T) {
		t.Run("en", func(t *testing.T) {
			t.Run("len=4, n=3", func(t *testing.T) {
				result, err := From("word", 3)
				require.NoError(t, err)
				require.Equal(t, []string{"wor", "ord"}, result)
			})
			t.Run("len=4, n=2", func(t *testing.T) {
				result, err := From("word", 2)
				require.NoError(t, err)
				require.Equal(t, []string{"wo", "or", "rd"}, result)
			})
			t.Run("len=6, n=3", func(t *testing.T) {
				result, err := From("orange", 3)
				require.NoError(t, err)
				require.Equal(t, []string{"ora", "ran", "ang", "nge"}, result)
			})
		})

		t.Run("ru", func(t *testing.T) {
			t.Run("len=6, n=3", func(t *testing.T) {
				result, err := From("яблоко", 3)
				require.NoError(t, err)
				require.Equal(t, []string{"ябл", "бло", "лок", "око"}, result)
			})
		})
	})
}
