package env

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnvAsStringOrFallback(t *testing.T) {
	const want = "TEST_STRING_ENV"

	assert := assert.New(t)

	key := "TEST_STRING_KEY"
	t.Setenv(key, want)
	assert.Equal(want, GetEnvAsStringOrFallback(key, "~"+want))

	key = "TEST_STRING_KEY_NOT_EXIST"
	assert.Equal(want, GetEnvAsStringOrFallback(key, want))
}

func TestGetEnvAsIntOrFallback(t *testing.T) {
	const want = 123

	assert := assert.New(t)

	key := "TEST_INT_KEY"
	t.Setenv(key, strconv.Itoa(want))
	val, _ := GetEnvAsIntOrFallback(key, 123)
	assert.Equal(want, val)

	key = "TEST_INT_KEY_NOT_EXIST"
	val, _ = GetEnvAsIntOrFallback(key, want)
	assert.Equal(want, val)

	key = "TEST_INT_KEY"
	t.Setenv(key, "not int")
	val, err := GetEnvAsIntOrFallback(key, want)
	assert.Equal(want, val)
	if err == nil {
		t.Error("expected error")
	}
}

func TestGetEnvAsFloat64OrFallback(t *testing.T) {
	const want = 123.456

	assert := assert.New(t)

	key := "TEST_FLOAT_KEY"
	t.Setenv(key, "123.456")
	val, _ := GetEnvAsFloat64OrFallback(key, 123.456)
	assert.Equal(want, val)

	key = "TEST_FLOAT_KEY_NOT_EXIST"
	val, _ = GetEnvAsFloat64OrFallback(key, want)
	assert.Equal(want, val)

	key = "TEST_FLOAT_KEY"
	t.Setenv(key, "not float")
	val, err := GetEnvAsFloat64OrFallback(key, want)
	assert.Equal(want, val)
	if err == nil {
		t.Error("expected error")
	}
}
