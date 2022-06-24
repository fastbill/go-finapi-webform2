package webform2

import (
	"strconv"
	"time"
)

type Time time.Time

const RFC3339NoColon = "2006-01-02T15:04:05Z0700"

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

// UnixTime defines a timestamp encoded as epoch seconds in JSON
type UnixTime time.Time

// MarshalJSON is used to convert the timestamp to JSON
func (t UnixTime) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(time.Time(t).Unix(), 10)), nil
}

// UnmarshalJSON is used to convert the timestamp from JSON
func (t *UnixTime) UnmarshalJSON(s []byte) (err error) {
	r := string(s)
	q, err := strconv.ParseInt(r, 10, 64)
	if err != nil {
		return err
	}
	*(*time.Time)(t) = time.Unix(q, 0)
	return nil
}
