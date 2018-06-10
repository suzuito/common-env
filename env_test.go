package env

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnv(t *testing.T) {
	os.Setenv("a", "hoge")
	os.Setenv("b", "1")
	os.Setenv("c", "true")
	assert.Equal(t, "hoge", GetenvAsString("a", "foo"))
	assert.Equal(t, "foo", GetenvAsString("aa", "foo"))
	assert.Equal(t, "foo", GetenvAsString("", "foo"))
	assert.Equal(t, 1, GetenvAsInt("b", 2))
	assert.Equal(t, 2, GetenvAsInt("bb", 2))
	assert.Equal(t, 2, GetenvAsInt("a", 2))
	assert.Equal(t, 2, GetenvAsInt("", 2))
	assert.Equal(t, true, GetenvAsBool("c", false))
	assert.Equal(t, false, GetenvAsBool("cc", false))
	assert.Equal(t, false, GetenvAsBool("a", false))
	assert.Equal(t, false, GetenvAsBool("", false))
}
