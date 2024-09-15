package envx

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnvStrWhenEnvIsSet(t *testing.T) {
	key := "TEST_ENV_STR"
	expected := "HelloWorld"
	os.Setenv(key, expected)

	result := GetEnvStr(key, "DefaultValue")

	assert.Equal(t, expected, result)

	os.Unsetenv(key)
}

func TestGetEnvStrWhenEnvIsNotSet(t *testing.T) {
	key := "TEST_ENV_STR"
	defaultValue := "DefaultValue"

	result := GetEnvStr(key, defaultValue)

	assert.Equal(t, defaultValue, result)
}

func TestGetEnvIntWhenEnvIsSet(t *testing.T) {
	key := "TEST_ENV_INT"
	expected := 42
	os.Setenv(key, "42")

	result := GetEnvInt(key, 100)

	assert.Equal(t, expected, result)

	os.Unsetenv(key)
}

func TestGetEnvIntWhenEnvIsNotSet(t *testing.T) {
	key := "TEST_ENV_INT"
	defaultValue := 100

	result := GetEnvInt(key, defaultValue)

	assert.Equal(t, defaultValue, result)
}

func TestGetEnvIntWhenEnvIsInvalid(t *testing.T) {
	key := "TEST_ENV_INT"
	defaultValue := 100
	os.Setenv(key, "invalid")

	result := GetEnvInt(key, defaultValue)

	assert.Equal(t, defaultValue, result)

	os.Unsetenv(key)
}

func TestGetEnvInt64WhenEnvIsSet(t *testing.T) {
	key := "TEST_ENV_INT64"
	expected := int64(64)
	os.Setenv(key, "64")

	result := GetEnvInt64(key, 100)

	assert.Equal(t, expected, result)

	os.Unsetenv(key)
}

func TestGetEnvInt64WhenEnvIsNotSet(t *testing.T) {
	key := "TEST_ENV_INT64"
	defaultValue := int64(100)

	result := GetEnvInt64(key, defaultValue)

	assert.Equal(t, defaultValue, result)
}

func TestGetEnvBoolWhenEnvIsSetToTrue(t *testing.T) {
	key := "TEST_ENV_BOOL"
	os.Setenv(key, "true")

	result := GetEnvBool(key, false)

	assert.True(t, result)

	os.Unsetenv(key)
}

func TestGetEnvBoolWhenEnvIsSetToFalse(t *testing.T) {
	key := "TEST_ENV_BOOL"
	os.Setenv(key, "false")

	result := GetEnvBool(key, true)

	assert.False(t, result)

	os.Unsetenv(key)
}

func TestGetEnvBoolWhenEnvIsNotSet(t *testing.T) {
	key := "TEST_ENV_BOOL"
	defaultValue := true

	result := GetEnvBool(key, defaultValue)

	assert.Equal(t, defaultValue, result)
}
