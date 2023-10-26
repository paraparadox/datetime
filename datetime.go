package datetime

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

type Date time.Time

const DateLayout = "2006-01-02"

func (d *Date) Scan(value interface{}) (err error) {
	nullTime := &sql.NullTime{}
	err = nullTime.Scan(value)
	*d = Date(nullTime.Time)
	return
}

func (d Date) Value() (driver.Value, error) {
	year, month, day := time.Time(d).Date()
	return time.Date(year, month, day, 0, 0, 0, 0, time.Time(d).Location()), nil
}

// GormDataType gorm common data type
func (d Date) GormDataType() string {
	return "date"
}

func (d Date) GobEncode() ([]byte, error) {
	return time.Time(d).GobEncode()
}

func (d *Date) GobDecode(b []byte) error {
	return (*time.Time)(d).GobDecode(b)
}

func (d Date) MarshalJSON() ([]byte, error) {
	if time.Time(d).IsZero() {
		return []byte(`""`), nil
	}

	return []byte(fmt.Sprintf(`"%s"`, time.Time(d).Format(DateLayout))), nil
}

func (d *Date) UnmarshalJSON(input []byte) error {
	s := strings.Trim(string(input), `"`)
	parsedDate, err := time.Parse(DateLayout, s)
	*d = Date(parsedDate)
	return err
}

func (d Date) String() string {
	return time.Time(d).Format(DateLayout)
}

type Time time.Time

const TimeLayout = "15:04"

func (t *Time) Scan(value interface{}) (err error) {
	nullTime := &sql.NullTime{}
	err = nullTime.Scan(value)
	*t = Time(nullTime.Time)
	return
}

func (t Time) Value() (driver.Value, error) {
	hour := time.Time(t).Hour()
	minute := time.Time(t).Minute()
	return time.Date(1, 1, 1, hour, minute, 0, 0, time.Time(t).Location()), nil
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
		return []byte(`""`), nil
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

type DateTime time.Time

const DateTimeLayout = time.RFC3339Nano

func (d *DateTime) Scan(value interface{}) (err error) {
	nullTime := &sql.NullTime{}
	err = nullTime.Scan(value)
	*d = DateTime(nullTime.Time)
	return
}

func (d DateTime) Value() (driver.Value, error) {
	return time.Time(d), nil
}

// GormDataType gorm common data type
func (d DateTime) GormDataType() string {
	return "timestamp with time zone"
}

func (d DateTime) GobEncode() ([]byte, error) {
	return time.Time(d).GobEncode()
}

func (d *DateTime) GobDecode(b []byte) error {
	return (*time.Time)(d).GobDecode(b)
}

func (d DateTime) MarshalJSON() ([]byte, error) {
	if time.Time(d).IsZero() {
		return []byte(`""`), nil
	}

	return []byte(fmt.Sprintf(`"%s"`, time.Time(d).Format(DateTimeLayout))), nil
}

func (d *DateTime) UnmarshalJSON(input []byte) error {
	s := strings.Trim(string(input), `"`)
	parsedDateTime, err := time.Parse(DateTimeLayout, s)
	*d = DateTime(parsedDateTime)
	return err
}

func (d DateTime) String() string {
	return time.Time(d).Format(DateTimeLayout)
}
