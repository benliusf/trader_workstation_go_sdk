package client

import (
	"strings"

	api "github.com/benliusf/trader_workstation_go_sdk/api/v104401"
)

func IsPaperTrading(a *api.AccountSummary) bool {
	return a != nil && a.Account != nil && strings.HasPrefix(*a.Account, "DU")
}
