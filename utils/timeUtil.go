package utils

import "time"

const TIME_LAYOUT = "2006-01-02 15:04:05"

func DatatimeToUnix(date string) int64 {
	t, _ := time.Parse(TIME_LAYOUT, date)
	tLocal := t.Local()

	return tLocal.Unix()
}

func DatetimeToTime(date string) time.Time {
	t, _ := time.Parse(TIME_LAYOUT, date)
	tLocal := t.Local()

	return tLocal
}

func UnixToString(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format(TIME_LAYOUT)
}

func UnixToTime(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}
