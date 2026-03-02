package timex

import "testing"
import "time"

func withLocalLocation(t *testing.T, loc *time.Location) {
	t.Helper()

	original := time.Local
	time.Local = loc
	t.Cleanup(func() {
		time.Local = original
	})
}

func TestParseAndParseInLocation(t *testing.T) {
	local := time.FixedZone("LOCAL+8", 8*3600)
	withLocalLocation(t, local)

	got, err := Parse(LayoutDateTime, "2026-03-02 10:20:30")
	if err != nil {
		t.Fatalf("Parse returned unexpected error: %v", err)
	}
	want := time.Date(2026, 3, 2, 10, 20, 30, 0, local)
	if !got.Equal(want) {
		t.Fatalf("Parse mismatch: got=%v want=%v", got, want)
	}
	if got.Location() != local {
		t.Fatalf("Parse location mismatch: got=%v want=%v", got.Location(), local)
	}

	gotNilLoc, err := ParseInLocation(LayoutDate, "2026-03-02", nil)
	if err != nil {
		t.Fatalf("ParseInLocation(nil) returned unexpected error: %v", err)
	}
	wantNilLoc := time.Date(2026, 3, 2, 0, 0, 0, 0, local)
	if !gotNilLoc.Equal(wantNilLoc) {
		t.Fatalf("ParseInLocation(nil) mismatch: got=%v want=%v", gotNilLoc, wantNilLoc)
	}
	if gotNilLoc.Location() != local {
		t.Fatalf("ParseInLocation(nil) location mismatch: got=%v want=%v", gotNilLoc.Location(), local)
	}

	shanghai := time.FixedZone("CST+8", 8*3600)
	gotWithLoc, err := ParseInLocation(LayoutDateTime, "2026-03-02 12:30:40", shanghai)
	if err != nil {
		t.Fatalf("ParseInLocation returned unexpected error: %v", err)
	}
	wantWithLoc := time.Date(2026, 3, 2, 12, 30, 40, 0, shanghai)
	if !gotWithLoc.Equal(wantWithLoc) {
		t.Fatalf("ParseInLocation mismatch: got=%v want=%v", gotWithLoc, wantWithLoc)
	}
	if gotWithLoc.Location() != shanghai {
		t.Fatalf("ParseInLocation location mismatch: got=%v want=%v", gotWithLoc.Location(), shanghai)
	}
}

func TestParseAndHelpersErrors(t *testing.T) {
	testCases := []struct {
		name string
		fn   func() (time.Time, error)
	}{
		{
			name: "Parse",
			fn: func() (time.Time, error) {
				return Parse(LayoutDateTime, "bad-value")
			},
		},
		{
			name: "ParseDate",
			fn: func() (time.Time, error) {
				return ParseDate("bad-date")
			},
		},
		{
			name: "ParseDateTime",
			fn: func() (time.Time, error) {
				return ParseDateTime("bad-datetime")
			},
		},
		{
			name: "ParseRFC3339",
			fn: func() (time.Time, error) {
				return ParseRFC3339("bad-rfc3339")
			},
		},
	}

	for _, tc := range testCases {
		_, err := tc.fn()
		if err == nil {
			t.Fatalf("%s expected error, got nil", tc.name)
		}
	}
}

func assertPanic(t *testing.T, name string, fn func()) {
	t.Helper()

	defer func() {
		if recover() == nil {
			t.Fatalf("%s expected panic, got none", name)
		}
	}()
	fn()
}

func TestMustFunctionsPanic(t *testing.T) {
	local := time.FixedZone("LOCAL+8", 8*3600)
	withLocalLocation(t, local)

	assertPanic(t, "MustParse", func() {
		_ = MustParse(LayoutDate, "bad")
	})
	assertPanic(t, "MustParseInLocation", func() {
		_ = MustParseInLocation(LayoutDateTime, "bad", local)
	})
	assertPanic(t, "MustParseDate", func() {
		_ = MustParseDate("bad")
	})
	assertPanic(t, "MustParseDateTime", func() {
		_ = MustParseDateTime("bad")
	})
	assertPanic(t, "MustParseRFC3339", func() {
		_ = MustParseRFC3339("bad")
	})
}

func TestHelperParseAndFormat(t *testing.T) {
	local := time.FixedZone("LOCAL+8", 8*3600)
	withLocalLocation(t, local)

	gotDate, err := ParseDate("2026-03-02")
	if err != nil {
		t.Fatalf("ParseDate returned unexpected error: %v", err)
	}
	wantDate := time.Date(2026, 3, 2, 0, 0, 0, 0, local)
	if !gotDate.Equal(wantDate) {
		t.Fatalf("ParseDate mismatch: got=%v want=%v", gotDate, wantDate)
	}
	if gotDate.Location() != local {
		t.Fatalf("ParseDate location mismatch: got=%v want=%v", gotDate.Location(), local)
	}

	gotDateTime, err := ParseDateTime("2026-03-02 11:22:33")
	if err != nil {
		t.Fatalf("ParseDateTime returned unexpected error: %v", err)
	}
	wantDateTime := time.Date(2026, 3, 2, 11, 22, 33, 0, local)
	if !gotDateTime.Equal(wantDateTime) {
		t.Fatalf("ParseDateTime mismatch: got=%v want=%v", gotDateTime, wantDateTime)
	}
	if gotDateTime.Location() != local {
		t.Fatalf("ParseDateTime location mismatch: got=%v want=%v", gotDateTime.Location(), local)
	}

	rfcInput := "2026-03-02T09:10:11+09:00"
	gotRFC3339, err := ParseRFC3339(rfcInput)
	if err != nil {
		t.Fatalf("ParseRFC3339 returned unexpected error: %v", err)
	}
	if FormatRFC3339(gotRFC3339) != rfcInput {
		t.Fatalf("ParseRFC3339/FormatRFC3339 mismatch: got=%s want=%s", FormatRFC3339(gotRFC3339), rfcInput)
	}
	if gotRFC3339.Location() == local {
		t.Fatalf("ParseRFC3339 should not force local location")
	}

	if Format(LayoutDate, gotDateTime) != "2026-03-02" {
		t.Fatalf("Format mismatch")
	}
	if FormatDate(gotDateTime) != "2026-03-02" {
		t.Fatalf("FormatDate mismatch")
	}
	if FormatDateTime(gotDateTime) != "2026-03-02 11:22:33" {
		t.Fatalf("FormatDateTime mismatch")
	}
}

func TestUnixConversions(t *testing.T) {
	local := time.FixedZone("LOCAL+5", 5*3600)
	withLocalLocation(t, local)

	ref := time.Unix(1700000000, 123000000).In(local)
	if ToUnix(ref) != 1700000000 {
		t.Fatalf("ToUnix mismatch: got=%d want=%d", ToUnix(ref), int64(1700000000))
	}
	if ToUnixMilli(ref) != 1700000000123 {
		t.Fatalf("ToUnixMilli mismatch: got=%d want=%d", ToUnixMilli(ref), int64(1700000000123))
	}

	fromSec := FromUnix(1700000000)
	if fromSec.Unix() != 1700000000 {
		t.Fatalf("FromUnix unix mismatch: got=%d want=%d", fromSec.Unix(), int64(1700000000))
	}
	if fromSec.Location() != local {
		t.Fatalf("FromUnix location mismatch: got=%v want=%v", fromSec.Location(), local)
	}

	fromMilli := FromUnixMilli(1700000000123)
	if fromMilli.UnixMilli() != 1700000000123 {
		t.Fatalf("FromUnixMilli unix milli mismatch: got=%d want=%d", fromMilli.UnixMilli(), int64(1700000000123))
	}
	if fromMilli.Location() != local {
		t.Fatalf("FromUnixMilli location mismatch: got=%v want=%v", fromMilli.Location(), local)
	}

	nowSec := NowUnix()
	nowMilli := NowUnixMilli()
	if nowMilli/1000 < nowSec-1 || nowMilli/1000 > nowSec+1 {
		t.Fatalf("NowUnix and NowUnixMilli inconsistent: sec=%d milli=%d", nowSec, nowMilli)
	}
}

func TestStartAndEndOfDay(t *testing.T) {
	loc := time.FixedZone("LOCAL+8", 8*3600)
	input := time.Date(2026, 3, 2, 15, 4, 5, 987654321, loc)

	start := StartOfDay(input)
	end := EndOfDay(input)

	wantStart := time.Date(2026, 3, 2, 0, 0, 0, 0, loc)
	wantEnd := wantStart.AddDate(0, 0, 1).Add(-time.Nanosecond)

	if !start.Equal(wantStart) {
		t.Fatalf("StartOfDay mismatch: got=%v want=%v", start, wantStart)
	}
	if !end.Equal(wantEnd) {
		t.Fatalf("EndOfDay mismatch: got=%v want=%v", end, wantEnd)
	}
	if start.Location() != loc || end.Location() != loc {
		t.Fatalf("day boundary location mismatch: start=%v end=%v want=%v", start.Location(), end.Location(), loc)
	}
	if !end.Equal(StartOfDay(input).AddDate(0, 0, 1).Add(-time.Nanosecond)) {
		t.Fatalf("EndOfDay relation mismatch")
	}
}

func TestStartAndEndOfWeekMondayRule(t *testing.T) {
	loc := time.FixedZone("LOCAL+8", 8*3600)
	tuesday := time.Date(2026, 3, 3, 10, 0, 0, 0, loc)
	sunday := time.Date(2026, 3, 8, 21, 30, 0, 0, loc)

	wantStart := time.Date(2026, 3, 2, 0, 0, 0, 0, loc)
	gotStartTuesday := StartOfWeek(tuesday)
	gotStartSunday := StartOfWeek(sunday)

	if !gotStartTuesday.Equal(wantStart) {
		t.Fatalf("StartOfWeek(tuesday) mismatch: got=%v want=%v", gotStartTuesday, wantStart)
	}
	if !gotStartSunday.Equal(wantStart) {
		t.Fatalf("StartOfWeek(sunday) mismatch: got=%v want=%v", gotStartSunday, wantStart)
	}

	wantEnd := wantStart.AddDate(0, 0, 7).Add(-time.Nanosecond)
	gotEnd := EndOfWeek(tuesday)
	if !gotEnd.Equal(wantEnd) {
		t.Fatalf("EndOfWeek mismatch: got=%v want=%v", gotEnd, wantEnd)
	}
	if gotStartTuesday.Location() != loc || gotEnd.Location() != loc {
		t.Fatalf("week boundary location mismatch: start=%v end=%v want=%v", gotStartTuesday.Location(), gotEnd.Location(), loc)
	}
	if !gotEnd.Equal(StartOfWeek(tuesday).AddDate(0, 0, 7).Add(-time.Nanosecond)) {
		t.Fatalf("EndOfWeek relation mismatch")
	}
}

func TestStartAndEndOfMonthLeapYear(t *testing.T) {
	loc := time.FixedZone("LOCAL+8", 8*3600)
	input := time.Date(2024, 2, 20, 12, 0, 0, 0, loc)

	start := StartOfMonth(input)
	end := EndOfMonth(input)

	wantStart := time.Date(2024, 2, 1, 0, 0, 0, 0, loc)
	wantEnd := time.Date(2024, 2, 29, 23, 59, 59, int(time.Second-time.Nanosecond), loc)

	if !start.Equal(wantStart) {
		t.Fatalf("StartOfMonth mismatch: got=%v want=%v", start, wantStart)
	}
	if !end.Equal(wantEnd) {
		t.Fatalf("EndOfMonth mismatch: got=%v want=%v", end, wantEnd)
	}
	if start.Location() != loc || end.Location() != loc {
		t.Fatalf("month boundary location mismatch: start=%v end=%v want=%v", start.Location(), end.Location(), loc)
	}
	if !end.Equal(StartOfMonth(input).AddDate(0, 1, 0).Add(-time.Nanosecond)) {
		t.Fatalf("EndOfMonth relation mismatch")
	}
}
