package timex

import "time"

const (
	LayoutDate     = "2006-01-02"
	LayoutDateTime = "2006-01-02 15:04:05"
	LayoutRFC3339  = time.RFC3339
)

func Parse(layout, value string) (time.Time, error) {
	return ParseInLocation(layout, value, time.Local)
}

func MustParse(layout, value string) time.Time {
	t, err := Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return t
}

func ParseInLocation(layout, value string, loc *time.Location) (time.Time, error) {
	if loc == nil {
		loc = time.Local
	}
	return time.ParseInLocation(layout, value, loc)
}

func MustParseInLocation(layout, value string, loc *time.Location) time.Time {
	t, err := ParseInLocation(layout, value, loc)
	if err != nil {
		panic(err)
	}
	return t
}

func ParseDate(value string) (time.Time, error) {
	return Parse(LayoutDate, value)
}

func MustParseDate(value string) time.Time {
	t, err := ParseDate(value)
	if err != nil {
		panic(err)
	}
	return t
}

func ParseDateTime(value string) (time.Time, error) {
	return Parse(LayoutDateTime, value)
}

func MustParseDateTime(value string) time.Time {
	t, err := ParseDateTime(value)
	if err != nil {
		panic(err)
	}
	return t
}

func ParseRFC3339(value string) (time.Time, error) {
	return time.Parse(time.RFC3339, value)
}

func MustParseRFC3339(value string) time.Time {
	t, err := ParseRFC3339(value)
	if err != nil {
		panic(err)
	}
	return t
}

func Format(layout string, t time.Time) string {
	return t.Format(layout)
}

func FormatDate(t time.Time) string {
	return Format(LayoutDate, t)
}

func FormatDateTime(t time.Time) string {
	return Format(LayoutDateTime, t)
}

func FormatRFC3339(t time.Time) string {
	return Format(LayoutRFC3339, t)
}

func ToUnix(t time.Time) int64 {
	return t.Unix()
}

func ToUnixMilli(t time.Time) int64 {
	return t.UnixMilli()
}

func FromUnix(sec int64) time.Time {
	return time.Unix(sec, 0).In(time.Local)
}

func FromUnixMilli(ms int64) time.Time {
	return time.UnixMilli(ms).In(time.Local)
}

func NowUnix() int64 {
	return time.Now().Unix()
}

func NowUnixMilli() int64 {
	return time.Now().UnixMilli()
}

func StartOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

func EndOfDay(t time.Time) time.Time {
	return StartOfDay(t).AddDate(0, 0, 1).Add(-time.Nanosecond)
}

func StartOfWeek(t time.Time) time.Time {
	day := int(t.Weekday())
	if day == int(time.Sunday) {
		day = 7
	}
	return StartOfDay(t).AddDate(0, 0, -(day - 1))
}

func EndOfWeek(t time.Time) time.Time {
	return StartOfWeek(t).AddDate(0, 0, 7).Add(-time.Nanosecond)
}

func StartOfMonth(t time.Time) time.Time {
	year, month, _ := t.Date()
	return time.Date(year, month, 1, 0, 0, 0, 0, t.Location())
}

func EndOfMonth(t time.Time) time.Time {
	return StartOfMonth(t).AddDate(0, 1, 0).Add(-time.Nanosecond)
}
