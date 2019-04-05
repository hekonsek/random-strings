package randomstrings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestForHuman(t *testing.T) {
	t.Parallel()
	assert.NotEmpty(t, ForHuman())
}

func TestForHumanWithHash(t *testing.T) {
	t.Parallel()

	name := ForHumanWithHash()
	assert.NotEmpty(t, name)
}

func TestForHumanMustBeLongerThanEightCharacters(t *testing.T) {
	t.Parallel()

	name := ForHuman()
	assert.True(t, len(name) >= 8)
}

func TestForHumanWithDashAndHash(t *testing.T) {
	t.Parallel()
	assert.NotEmpty(t, ForHumanWithDashAndHash())
}

func TestForHumanWithDashAndHashContainsDash(t *testing.T) {
	t.Parallel()
	assert.Contains(t, ForHumanWithDashAndHash(), "-")
}
