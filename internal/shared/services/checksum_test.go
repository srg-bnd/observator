package services

import (
	"fmt"
	"testing"

	"github.com/srg-bnd/observator/tests/helpers"
	"github.com/stretchr/testify/assert"
)

func TestChecksum(t *testing.T) {
	randomString, err := helpers.SecureRandomString(16)
	assert.Nil(t, err)

	testCases := []struct {
		data     string
		dataHash string
		err      error
	}{
		{
			data:     "correctData",
			dataHash: helpers.Sha256Hash("correctData", randomString),
		},
		{
			data:     "",
			dataHash: helpers.Sha256Hash("", randomString),
		},
		{
			data:     "correctData",
			dataHash: helpers.Sha256Hash("incorrectData", randomString),
			err:      ErrVerify,
		},
	}

	service := NewChecksum(randomString)

	for i, tc := range testCases {
		t.Run(fmt.Sprint("Test ", i+1, ": ", tc.data), func(t *testing.T) {
			err := service.Verify(tc.dataHash, tc.data)

			if tc.err != nil {
				assert.NotNil(t, err)
			} else {
				assert.ErrorIs(t, err, tc.err)
			}
		})
	}
}

func TestSum(t *testing.T) {
	randomString, err := helpers.SecureRandomString(16)
	assert.Nil(t, err)

	testCases := []struct {
		input    string
		expected string
		fail     bool
	}{
		{
			input:    "correctData",
			expected: helpers.Sha256Hash("correctData", randomString),
		},
		// For testing resets
		{
			input:    "newCorrectData",
			expected: helpers.Sha256Hash("newCorrectData", randomString),
		},
		{
			input:    "oldData",
			expected: helpers.Sha256Hash("newData", randomString),
			fail:     true,
		},
	}

	service := NewChecksum(randomString)

	for i, tc := range testCases {
		t.Run(fmt.Sprint("Test ", i+1, ": ", tc.input), func(t *testing.T) {
			result, err := service.Sum(tc.input)
			assert.Nil(t, err)

			if tc.fail {
				assert.NotEqual(t, tc.expected, result)
			} else {
				assert.Equal(t, tc.expected, result)
			}
		})
	}
}
