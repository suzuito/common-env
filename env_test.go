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
	os.Setenv("d", "1.1")
	assert.Equal(t, "hoge", GetenvAsString("a", "foo"))
	assert.Equal(t, "foo", GetenvAsString("aa", "foo"))
	assert.Equal(t, "foo", GetenvAsString("", "foo"))
	assert.Equal(t, 1, GetenvAsInt("b", 2))
	assert.Equal(t, 2, GetenvAsInt("bb", 2))
	assert.Equal(t, 2, GetenvAsInt("a", 2))
	assert.Equal(t, 2, GetenvAsInt("", 2))
	assert.Equal(t, int64(1), GetenvAsInt64("b", 2))
	assert.Equal(t, int64(2), GetenvAsInt64("bb", 2))
	assert.Equal(t, int64(2), GetenvAsInt64("a", 2))
	assert.Equal(t, int64(2), GetenvAsInt64("", 2))
	assert.Equal(t, true, GetenvAsBool("c", false))
	assert.Equal(t, false, GetenvAsBool("cc", false))
	assert.Equal(t, false, GetenvAsBool("a", false))
	assert.Equal(t, false, GetenvAsBool("", false))
	assert.Equal(t, float32(1.1), GetenvAsFloat32("d", 1.2))
	assert.Equal(t, float32(1.2), GetenvAsFloat32("dd", 1.2))
	assert.Equal(t, float32(1.2), GetenvAsFloat32("a", 1.2))
	assert.Equal(t, float32(1.2), GetenvAsFloat32("", 1.2))
	assert.Equal(t, float64(1.1), GetenvAsFloat64("d", 1.2))
	assert.Equal(t, float64(1.2), GetenvAsFloat64("dd", 1.2))
	assert.Equal(t, float64(1.2), GetenvAsFloat64("a", 1.2))
	assert.Equal(t, float64(1.2), GetenvAsFloat64("", 1.2))
}
