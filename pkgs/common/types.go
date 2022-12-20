package common

import (
	"encoding/json"
	"time"
)

type DateTime time.Time

var _ json.Unmarshaler = &DateTime{}

const (
	timeFormat = "2006-01-02 15:04:05"
)

func (mt *DateTime) UnmarshalJSON(bs []byte) error {
	var s string
	err := json.Unmarshal(bs, &s)
	if err != nil {
		return err
	}
	t, err := time.ParseInLocation(timeFormat, s, time.UTC)
	if err != nil {
		return err
	}
	*mt = DateTime(t)
	return nil
}

func (t DateTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormat)
	b = append(b, '"')
	return b, nil
}
