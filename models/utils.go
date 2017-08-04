package models

import (
	"time"
)

const (
	defaultTimeLayout = "Mon Jan 02 15:04:05 -0700 2006"
)

func StringToTime(str string) *time.Time {
	t, err := time.Parse(defaultTimeLayout, str)
	if err != nil {
		return nil
	}

	utc := t.UTC()
	return &utc
}

func TimeToString(t time.Time) string {
	return t.Format(defaultTimeLayout)
}