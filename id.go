package snowflake

import (
	"math"
	"time"
)

type IDInterface interface {
	ConvertToInteger() int64
	ExtractDatetime(startTimeFrom int64, timezone string) (*time.Time, error)
}

type ID struct {
	sequence  int16
	machineId int16
	timestamp int64
}

func (id *ID) ConvertToInteger() int64 {
	result := int64(id.sequence)
	result += int64(id.machineId) * int64(math.Pow(2, 12))
	result += id.timestamp * int64(math.Pow(2, 22))
	return result
}

func (id *ID) ExtractDatetime(startTimeFrom int64, timezone string) (*time.Time, error) {
	location, err := time.LoadLocation(timezone)
	if err != nil {
		return nil, err
	}
	currentTime := startTimeFrom + id.timestamp
	result := time.Unix(currentTime, 0).In(location)
	return &result, nil
}
