package models

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

type Time time.Time

const TimeLayout = "15:04"

func (t *Time) Scan(value interface{}) (err error) {
	nullTime := &sql.NullTime{}
	err = nullTime.Scan(value)
	*t = Time(nullTime.Time)
	return
}

func (t Time) Value() (driver.Value, error) {
	year, month, day := time.Time(t).Date()
	return time.Date(year, month, day, 0, 0, 0, 0, time.Time(t).Location()), nil
}

// GormDataType gorm common data type
func (t Time) GormDataType() string {
	return "time"
}

func (t Time) GobEncode() ([]byte, error) {
	return time.Time(t).GobEncode()
}

func (t *Time) GobDecode(b []byte) error {
	return (*time.Time)(t).GobDecode(b)
}

func (t Time) MarshalJSON() ([]byte, error) {
	if time.Time(t).IsZero() {
		return nil, nil
	}
	return []byte(fmt.Sprintf(`"%s"`, time.Time(t).Format(TimeLayout))), nil
}

func (t *Time) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	parsedTime, err := time.Parse(TimeLayout, s)
	*t = Time(parsedTime)
	return err
}

func (t Time) String() string {
	return time.Time(t).Format(TimeLayout)
}
