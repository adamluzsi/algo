package sorts_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type TestCase struct {
	Expected Ints
	Actually Ints
}

func SharedSpec(t *testing.T, subject func(Ints)) {
	t.Run("given the current sorting algorithm", func(t *testing.T) {
		t.Run("when collection is single element", func(t *testing.T) {
			t.Parallel()

			slice := Ints{1}
			subject(slice)
			require.Equal(t, Ints{1}, slice)
		})

		t.Run("when collection has multiple elements", func(t *testing.T) {

			t.Parallel()

			cases := []TestCase{
				TestCase{Actually: Ints{3, 2, 1}, Expected: Ints{1, 2, 3}},
				TestCase{Actually: Ints{4, 2, 2}, Expected: Ints{2, 2, 4}},
				TestCase{Actually: Ints{99, 3, 42, 1}, Expected: Ints{1, 3, 42, 99}},
				TestCase{Actually: Ints{98, 1, 2, 99, 3}, Expected: Ints{1, 2, 3, 98, 99}},
			}

			for _, testCase := range cases {
				subject(testCase.Actually)

				require.Equal(t, testCase.Expected, testCase.Actually)
			}

		})
	})
}

type Ints []int

func (ints Ints) Len() int {
	return len(ints)
}

func (ints Ints) Less(i int, j int) bool {
	return ints[i] < ints[j]
}

func (ints Ints) Swap(i int, j int) {
	ints[i], ints[j] = ints[j], ints[i]
}
