package prettytime_test

import (
	"testing"
	"time"

	. "github.com/andanhm/go-prettytime"
)

func TestFormat(t *testing.T) {
	tests := []struct {
		name string
		t    time.Time
		want string
	}{
		{name: "Just now", t: time.Now(), want: "just now"},
		{name: "Second", t: time.Now().Add(
			time.Hour*time.Duration(0) +
				time.Minute*time.Duration(0) +
				time.Second*time.Duration(1)),
			want: "1 second from now",
		},
		{name: "SecondAgo", t: time.Now().Add(
			time.Hour*time.Duration(0) +
				time.Minute*time.Duration(0) +
				time.Second*time.Duration(-1)),
			want: "1 second ago"},
		{name: "SecondsAgo", t: time.Now().Add(
			time.Hour*time.Duration(0) +
				time.Minute*time.Duration(0) +
				time.Second*time.Duration(-2)),
			want: "2 seconds ago"},
		{name: "Minutes", t: time.Now().Add(time.Hour*time.Duration(0) +
			time.Minute*time.Duration(59) +
			time.Second*time.Duration(59)), want: "60 minutes from now"},
		{name: "Tomorrow", t: time.Now().AddDate(0, 0, 1), want: "tomorrow"},
		{name: "Yesterday", t: time.Now().AddDate(0, 0, -1), want: "yesterday"},
		{name: "Week", t: time.Now().AddDate(0, 0, 7), want: "1 week from now"},
		{name: "WeekAgo", t: time.Now().AddDate(0, 0, -7), want: "1 week ago"},
		{name: "Month", t: time.Now().AddDate(0, 1, 0), want: "1 month from now"},
		{name: "MonthAgo", t: time.Now().AddDate(0, -1, 0), want: "1 month ago"},
		{name: "Year", t: time.Now().AddDate(50, 0, 0), want: "50 years from now"},
		{name: "YearAgo", t: time.Now().AddDate(-2, 0, 0), want: "2 years ago"},
	}
	pretty := NewPrettyTimeFormatter("en-EN")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTimeSince := pretty.Format(tt.t); gotTimeSince != tt.want {
				t.Errorf("Format() = %v, want %v", gotTimeSince, tt.want)
			}
		})
	}
}

func TestFormatGerman(t *testing.T) {
	tests := []struct {
		name string
		t    time.Time
		want string
	}{
		{name: "Just now", t: time.Now(), want: "jetzt"},
		{name: "Second", t: time.Now().Add(
			time.Hour*time.Duration(0) +
				time.Minute*time.Duration(0) +
				time.Second*time.Duration(1)),
			want: "1 Sekunde ab jetzt",
		},
		{name: "SecondAgo", t: time.Now().Add(
			time.Hour*time.Duration(0) +
				time.Minute*time.Duration(0) +
				time.Second*time.Duration(-1)),
			want: "1 Sekunde zuvor"},
		{name: "SecondsAgo", t: time.Now().Add(
			time.Hour*time.Duration(0) +
				time.Minute*time.Duration(0) +
				time.Second*time.Duration(-2)),
			want: "2 Sekunden zuvor"},
		{name: "Minutes", t: time.Now().Add(time.Hour*time.Duration(0) +
			time.Minute*time.Duration(59) +
			time.Second*time.Duration(59)), want: "60 Minuten ab jetzt"},
		{name: "Tomorrow", t: time.Now().AddDate(0, 0, 1), want: "morgen"},
		{name: "Yesterday", t: time.Now().AddDate(0, 0, -1), want: "gestern"},
		{name: "Week", t: time.Now().AddDate(0, 0, 7), want: "1 Woche ab jetzt"},
		{name: "WeekAgo", t: time.Now().AddDate(0, 0, -7), want: "1 Woche zuvor"},
		{name: "Month", t: time.Now().AddDate(0, 1, 0), want: "1 Monat ab jetzt"},
		{name: "MonthAgo", t: time.Now().AddDate(0, -1, 0), want: "1 Monat zuvor"},
		{name: "Year", t: time.Now().AddDate(50, 0, 0), want: "50 Jahre ab jetzt"},
		{name: "YearAgo", t: time.Now().AddDate(-2, 0, 0), want: "2 Jahre zuvor"},
	}

	pretty := NewPrettyTimeFormatter("de-DE")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTimeSince := pretty.Format(tt.t); gotTimeSince != tt.want {
				t.Errorf("Format() = %v, want %v", gotTimeSince, tt.want)
			}
		})
	}
}

func TestFormatYear(t *testing.T) {
	now := time.Now()

	pretty := NewPrettyTimeFormatter("en-EN")
	oneYearFromNow := now.AddDate(1, 0, 0)
	gotTimeSince := pretty.Format(oneYearFromNow)
	expected := "1 year from now"
	if gotTimeSince != expected {
		t.Errorf("got %v, want %v", expected, gotTimeSince)
	}
}
