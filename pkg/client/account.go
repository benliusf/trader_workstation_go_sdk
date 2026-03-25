package client

import (
	"strings"
)

func IsPaperTrading(accountId string) bool {
	return accountId != "" && strings.HasPrefix(accountId, "DU")
}
