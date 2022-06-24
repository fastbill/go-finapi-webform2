package webform2

import (
	"time"
)

type Time time.Time

const RFC3339NoColon = "2006-01-02T15:04:05Z0700"

// UnmarshalJSON allows to parse the not-quite RCF3339 timestamp that finAPI sends into a time.Time value.
// Code adapted from https://cs.opensource.google/go/go/+/refs/tags/go1.18.3:src/time/time.go;l=1298;drc=90462dfc3aa99649de90bb587af56a9cb0214665.
func (t *Time) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" {
		return nil
	}
	parsedTime, err := time.Parse(`"`+RFC3339NoColon+`"`, string(data))
	if err != nil {
		return err
	}
	*t = Time(parsedTime)
	return nil
}

func (t Time) MarshalJSON() ([]byte, error) {
	return time.Time(t).MarshalJSON()
}
