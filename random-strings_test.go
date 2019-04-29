package randomstrings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestForHumanIsNotEmpty(t *testing.T) {
	t.Parallel()
	assert.NotEmpty(t, ForHuman())
}

func TestForHumanDoesNotContainSpaces(t *testing.T) {
	t.Parallel()
	assert.NotContains(t, ForHuman(), " ")
}

func TestForHumanWithHashNotEmpty(t *testing.T) {
	t.Parallel()
	assert.NotEmpty(t, ForHumanWithHash())
}

func TestForHumanWithHashDoesNotContainSpaces(t *testing.T) {
	t.Parallel()
	assert.NotContains(t, ForHumanWithHash(), " ")
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
