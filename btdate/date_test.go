package btdate

import (
	"testing"
	"time"
)

func TestMarshal(t *testing.T) {
	want := Date{Year: 2017, Month: time.January, Day: 3}

	dateBytes, err := want.MarshalText()
	if err != nil {
		t.Fatalf("unexpected err: %s", err)
	}
	if string(dateBytes) != "2017-01-03" {
		t.Fatalf("MarshalText: got %s, want 2017-01-03", dateBytes)
	}
}

func TestString(t *testing.T) {
	date := Date{Year: 2017, Month: time.January, Day: 3}
	if date.String() != "2017-01-03" {
		t.Errorf("Date.String: got %s, want 2017-01-03", date)
	}
}

func TestTime(t *testing.T) {
	want := Date{Year: 2017, Month: time.January, Day: 3}
	got := want.Time().Format("2006-01-02")
	if got != want.String() {
		t.Errorf("Date.Time: got %s, want %s", got, want)
	}
}

func TestToday(t *testing.T) {
	want := FromTime(time.Now().UTC())
	got := Today()
	if want.String() != got.String() {
		t.Errorf("Today: got %s, want %s", got, want)
	}

}

func TestUnmarshal(t *testing.T) {

	tests := []struct {
		name  string
		input []byte
		want  Date
	}{
		{
			name:  "normal",
			input: []byte("2017-01-03"),
			want:  Date{3, time.January, 2017},
		},
		{
			name:  "empty",
			input: nil,
			want:  Date{0, 0, 0},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			var got Date
			if err := got.UnmarshalText([]byte(test.input)); err != nil {
				t.Fatalf("unexpected err: %s", err)
			}
			if got != test.want {
				t.Errorf("UnmarshalText: got %s, want %s", got, test.want)
			}
		})
	}

	t.Run("malformed", func(t *testing.T) {
		var got Date
		if err := got.UnmarshalText([]byte("noDate")); err == nil {
			t.Errorf("expected err: parsing time..., but got nil")
		}
	})
}
