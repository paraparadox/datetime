package models

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
		return nil, nil
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
