package sorts_test

import (
	"testing"

	"github.com/adamluzsi/algo/sorts"
)

func TestMerge(t *testing.T) {
	SharedSpec(t, func(s Ints) { sorts.Merge(s) })
}
