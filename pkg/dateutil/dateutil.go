package dateutil

import "time"

const (
	DatetimeFormat    = "2006-01-02 15:04:05"
	DatetimeUTCFormat = "2006-01-02T15:04:05"
	DateFormat        = "2006-01-02"
	DateKeyFormat     = "20060102"
	TimeFormat        = "15:04:05"
)

func FormatDatetime(t time.Time) string {
	return t.Local().Format(DatetimeFormat)
}

func FormatDate(t time.Time) string {
	return t.Local().Format(DateFormat)
}

func FormatDateKey(t time.Time) string {
	return t.Local().Format(DateKeyFormat)
}

func FormatTime(t time.Time) string {
	return t.Local().Format(TimeFormat)
}

func FormatDatetimeUTC(t time.Time) string {
	return t.UTC().Format(DatetimeUTCFormat)
}
func FormatDateUTC(t time.Time) string {
	return t.UTC().Format(DateFormat)
}

func FormatDateKeyUTC(t time.Time) string {
	return t.UTC().Format(DateKeyFormat)
}

func FormatTimeUTC(t time.Time) string {
	return t.UTC().Format(TimeFormat)
}
