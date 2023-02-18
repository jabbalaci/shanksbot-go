package shanks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/////////////////////////////////////////////////////////////////////////////

func TestShanks(t *testing.T) {
	assert.Equal(t, Shanks(1), 0)
	assert.Equal(t, Shanks(2), 1)
	assert.Equal(t, Shanks(3), 1)
	assert.Equal(t, Shanks(5), 1)
	assert.Equal(t, Shanks(7), 6)
	assert.Equal(t, Shanks(23), 22)
	assert.Equal(t, Shanks(60_013), 5_001)
	assert.Equal(t, Shanks(60_017), 60_016)
	assert.Equal(t, Shanks(60_037), 10_006)
	assert.Equal(t, Shanks(61_561), 405)
}

// slow
// func TestSanity(t *testing.T) {
// for i := 50_000; i < 70_000; i++ {
// assert.Equal(t, Shanks(i) < i, true)
// }
// }
