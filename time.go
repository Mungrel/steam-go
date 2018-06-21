package steam

import (
	"fmt"
	"strconv"
	"time"
)

// Timestamp is an alias for time.Time
type Timestamp time.Time

// NewTimestamp is a convenience method for creating a *Timestamp
func NewTimestamp(unixTime int64) *Timestamp {
	stamp := Timestamp(time.Unix(unixTime, 0))
	return &stamp
}

// MarshalJSON marshals a Timestamp into JSON
func (t *Timestamp) MarshalJSON() ([]byte, error) {
	ts := time.Time(*t).Unix()
	stamp := fmt.Sprint(ts)

	return []byte(stamp), nil
}

// UnmarshalJSON unmarshals JSON into a Timestamp
func (t *Timestamp) UnmarshalJSON(b []byte) error {
	ts, err := strconv.Atoi(string(b))
	if err != nil {
		return err
	}

	*t = Timestamp(time.Unix(int64(ts), 0))

	return nil
}
