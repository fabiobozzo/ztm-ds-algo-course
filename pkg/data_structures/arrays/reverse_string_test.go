package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverse(t *testing.T) {
	assert.Equal(t, "oibaf onos oi oaic", reverse("ciao io sono fabio"))
	assert.Equal(t, "!oibaf onos oi oaic", reverse("ciao io sono fabio!"))
	assert.Equal(t, "a", reverse("a"))
	assert.Equal(t, "ba", reverse("ab"))
	assert.Equal(t, "", reverse(""))
	assert.Equal(t, "c b a _", reverse("_ a b c"))
}
