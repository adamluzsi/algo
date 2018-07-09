package sorts_test

import (
	"testing"

	"github.com/adamluzsi/algo/sorts"
)

func TestBubble(t *testing.T) {
	SharedSpec(t, func(is Ints) { sorts.Bubble(is) })
}
