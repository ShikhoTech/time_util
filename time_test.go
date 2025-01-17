package time_util

import (
	"reflect"
	"testing"
	"time"
)

func TestFromString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{"Test 1: UTC", args{"2021-08-01T00:00:00Z"}, time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC), false},
		{"Test 2: UTC+6", args{"2021-08-01T06:00:00+06:00"}, time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC), false},
		{"Test 3: UTC+6", args{"2021-08-01T00:00:00+06:00"}, time.Date(2021, 7, 31, 18, 0, 0, 0, time.UTC), false},
		{"Test 4: UTC-6", args{"2021-08-01T00:00:00-06:00"}, time.Date(2021, 8, 1, 6, 0, 0, 0, time.UTC), false},
		{"Test 5: Invalid Time", args{"2021"}, time.Time{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FromString(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromString() got = %v, want %v", got, tt.want)
			}
		})
	}
}
