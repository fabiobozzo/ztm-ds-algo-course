package recursion

import (
	"testing"
)

func TestTowerOfHanoi(t *testing.T) {
	tower := NewTowerOfHanoi(5)

	tower.Solve()
}
