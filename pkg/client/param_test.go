package client

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestQueryParams(t *testing.T) {
	now := time.Now()
	tests := []struct {
		startTime time.Time
		endTime   time.Time
		duration  string
	}{
		{
			startTime: now.Add(-1 * time.Hour),
			endTime:   now,
			duration:  "3600 S",
		},
		{
			startTime: now.Add(-48 * time.Hour),
			endTime:   now,
			duration:  "2 D",
		},
		{
			startTime: now.Add(-50 * time.Hour),
			endTime:   now,
			duration:  "3 D",
		},
		{
			startTime: now.Add(-168 * time.Hour),
			endTime:   now,
			duration:  "1 W",
		},
		{
			startTime: now.Add(-720 * time.Hour),
			endTime:   now,
			duration:  "1 M",
		},
	}
	for _, tt := range tests {
		q := &QueryParams{
			StartTime: tt.startTime,
			EndTime:   tt.endTime,
		}
		assert.Equal(t, tt.duration, q.Duration().String())
	}
}
