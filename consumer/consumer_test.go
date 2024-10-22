package consumer

import (
	"code-challenge/models"
	"testing"
)

func TestCheckForAnomalies(t *testing.T) {
	tests := []struct {
		data     models.Data
		expected bool
	}{
		{
			data:     models.Data{MeterReading: -1, MeterID: ""},
			expected: true,
		},
		{
			data:     models.Data{MeterReading: 1, MeterID: ""},
			expected: false,
		},
		{
			data:     models.Data{MeterReading: -1, MeterID: "meter1"},
			expected: false,
		},
		{
			data:     models.Data{MeterReading: 0, MeterID: ""},
			expected: false,
		},
		{
			data:     models.Data{MeterReading: 0, MeterID: "meter1"},
			expected: false,
		},
		{
			data:     models.Data{MeterReading: -5, MeterID: ""},
			expected: true,
		},
	}

	for _, test := range tests {
		result := checkForAnomalies(test.data)
		if result != test.expected {
			t.Errorf("checkForAnomalies(%v) = %v; expected %v", test.data, result, test.expected)
		}
	}
}
