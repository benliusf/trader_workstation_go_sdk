package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPaperTrading(t *testing.T) {
	tests := []struct {
		accountId string
		expected  bool
	}{
		{
			"", false,
		},
		{
			"123", false,
		},
		{
			"DU123", true,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.expected, IsPaperTrading(tt.accountId))
	}
}
