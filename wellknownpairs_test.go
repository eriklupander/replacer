package replacer

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAsBytePairs(t *testing.T) {
	out, err := AsBytePairs([]string{".", "", "$", "S"})
	require.NoError(t, err)
	assert.Equal(t, []byte{46, 8, 36, 83}, out)
}
