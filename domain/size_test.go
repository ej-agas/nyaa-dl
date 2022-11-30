package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSizeFromString(t *testing.T) {
	scenarios := []struct {
		testName      string
		input         string
		expectedBytes float64
		expectedErr   string
	}{
		{
			testName:      "Bytes",
			input:         "420 Bytes",
			expectedBytes: 420,
			expectedErr:   "",
		},
		{
			testName:      "Kibibyte (KiB)",
			input:         "1.2 KiB",
			expectedBytes: 1228.8,
			expectedErr:   "",
		},
		{
			testName:      "Mebibyte (MiB)",
			input:         "256.42 MiB",
			expectedBytes: 268_875_857.92,
			expectedErr:   "",
		},
		{
			testName:      "Gibibyte (GiB)",
			input:         "1014.24 GiB",
			expectedBytes: 1_089_031_907_573.76,
			expectedErr:   "",
		},
		{
			testName:      "Tebibyte (TiB)",
			input:         "421.24 TiB",
			expectedBytes: 463_158_278_084_362.25,
			expectedErr:   "",
		},
		{
			testName:      "Pebibyte (PiB)",
			input:         "1021.24 PiB",
			expectedBytes: 1_149_814_020_863_961_344,
			expectedErr:   "",
		},
		{
			testName:      "Exbibyte (EiB)",
			input:         "2047.42 EiB",
			expectedBytes: 2_360_514_546_962_150_719_488,
			expectedErr:   "",
		},
		{
			testName:      "Invalid Input",
			input:         "foobar",
			expectedBytes: 0.0,
			expectedErr:   "invalid size",
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.testName, func(t *testing.T) {
			size, err := ConvertToBytes(scenario.input)

			if scenario.expectedErr != "" {
				assert.Contains(t, err.Error(), scenario.expectedErr)
				return
			}

			assert.Equal(t, scenario.expectedBytes, size)
		})
	}
}
