package time_util

import (
	"context"
	"errors"
	"time"
)

// Now returns the current time in UTC
func Now() time.Time {
	return time.Now().UTC()
}

// ToString converts a time to a string in RFC3339 format. The time is first converted to UTC.
func ToString(t time.Time) string {
	return t.UTC().Format(time.RFC3339)
}

// FromString converts a string in RFC3339 format to a time. The time is returned in UTC.
func FromString(s string) (time.Time, error) {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return time.Time{}, errors.New("invalid time")
	}
	return t.UTC(), nil
}

func LocalCurrentTime(ctx context.Context) time.Time {
	location, err := time.LoadLocation(ctx.Value("tz_info").(string))
	if err != nil {
		return time.Now().UTC()
	}
	return time.Now().In(location)
}
