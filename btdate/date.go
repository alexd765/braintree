package btdate

import (
	"fmt"
	"time"
)

// Date is the braintree date.
//
// All dates are in UTC and other timezones are converted
// to and from UTC on all conversions.
type Date struct {
	Day   int
	Month time.Month
	Year  int
}

// FromTime converts time.Time into a Date.
func FromTime(t time.Time) Date {
	var date Date
	date.Year, date.Month, date.Day = t.UTC().Date()
	return date
}

// MarshalText implements the encoding.TextMarshaler interface.
func (d Date) MarshalText() ([]byte, error) {
	return []byte(d.String()), nil
}

// Parse a Date in the format YYYY-MM-DD.
func Parse(value string) (Date, error) {
	t, err := time.Parse("2006-01-02", value)
	if err != nil {
		return Date{}, err
	}
	return FromTime(t), nil
}

// String implements the Stringer interface.
func (d Date) String() string {
	return fmt.Sprintf("%d-%02d-%02d", d.Year, d.Month, d.Day)
}

// Time converts a Date into time.Time.
func (d Date) Time() time.Time {
	return time.Date(d.Year, d.Month, d.Day, 0, 0, 0, 0, time.UTC)
}

// Today returns the current day.
func Today() Date {
	return FromTime(time.Now())
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (d *Date) UnmarshalText(data []byte) error {
	date, err := Parse(string(data))
	if err != nil {
		return err
	}
	*d = date
	return nil
}
