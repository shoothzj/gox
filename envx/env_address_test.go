package envx

import (
	"github.com/shoothzj/gox/netx"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnvAddressListWhenEnvIsSet(t *testing.T) {
	key := "TEST_ENV_ADDRESSES"
	os.Setenv(key, "zookeeper-0:2181,zookeeper-1:2181")

	expected := []netx.Address{
		{Host: "zookeeper-0", Port: 2181},
		{Host: "zookeeper-1", Port: 2181},
	}

	result := GetEnvAddressList(key)

	assert.Equal(t, expected, result)

	os.Unsetenv(key)
}

func TestGetEnvAddressListWhenEnvIsNotSet(t *testing.T) {
	key := "TEST_ENV_ADDRESSES"

	result := GetEnvAddressList(key)

	assert.Empty(t, result)
}

func TestGetEnvAddressListWhenEnvHasInvalidFormat(t *testing.T) {
	key := "TEST_ENV_ADDRESSES"
	os.Setenv(key, "invalidAddress")

	// Expect an empty slice when the format is invalid
	result := GetEnvAddressList(key)

	assert.Empty(t, result)

	os.Unsetenv(key)
}

func TestGetEnvAddressWhenEnvIsSet(t *testing.T) {
	key := "TEST_ENV_ADDRESS"
	os.Setenv(key, "zookeeper-0:2181")

	expected := &netx.Address{Host: "zookeeper-0", Port: 2181}

	result := GetEnvAddress(key)

	assert.Equal(t, expected, result)

	os.Unsetenv(key)
}

func TestGetEnvAddressWhenEnvIsNotSet(t *testing.T) {
	key := "TEST_ENV_ADDRESS"

	result := GetEnvAddress(key)

	assert.Nil(t, result)
}

func TestGetEnvAddressWhenEnvHasInvalidFormat(t *testing.T) {
	key := "TEST_ENV_ADDRESS"
	os.Setenv(key, "invalidAddress")

	// Expect nil when the format is invalid
	result := GetEnvAddress(key)

	assert.Nil(t, result)

	os.Unsetenv(key)
}
