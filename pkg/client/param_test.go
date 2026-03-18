package client

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestQueryParams(t *testing.T) {
	now := time.Now()
	startTime := now.Add(-1 * time.Hour)
	q := &QueryParams{
		StartTime: startTime,
		EndTime:   now,
	}
	assert.Equal(t, "3600 S", q.Duration().String())
}
