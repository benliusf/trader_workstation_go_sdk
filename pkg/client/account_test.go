package client

import (
	"testing"

	api "github.com/benliusf/trader_workstation_go_sdk/api/v104401"
	"github.com/stretchr/testify/assert"
)

func TestIsPaperTrading(t *testing.T) {
	tests := []struct {
		acct     *api.AccountSummary
		expected bool
	}{
		{
			&api.AccountSummary{}, false,
		},
		{
			&api.AccountSummary{Account: strPtr("123")}, false,
		},
		{
			&api.AccountSummary{Account: strPtr("DU123")}, true,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.expected, IsPaperTrading(tt.acct))
	}
}
