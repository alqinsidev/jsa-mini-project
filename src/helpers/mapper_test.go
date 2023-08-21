package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapStatusDetail(t *testing.T) {
	tests := []struct {
		status   int64
		expected string
	}{
		{1, "Menunggu Verifikasi"},
		{2, "Terverifikasi"},
		{90, "Gagal Terverifikasi"},
		{100, "UNKNOWN STATUS"},
	}

	for _, test := range tests {
		result := MapStatusDetail(test.status)
		assert.Equal(t, test.expected, result)
	}
}
