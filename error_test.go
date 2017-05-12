package braintree

import "testing"
import "net/http"
import "bytes"
import "io/ioutil"

func TestAPIError(t *testing.T) {
	t.Parallel()

	err := &APIError{Code: 5, Message: "Johnny"}
	got := err.Error()
	want := "Code 5: Johnny"
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestProcessError(t *testing.T) {
	err := &ProcessorError{Code: 5, Message: "Johnny"}
	got := err.Error()
	want := "Code 5: Johnny"
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestParseError(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		resp      *http.Response
		wantError string
	}{
		{
			name: "unexpectedError",
			resp: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewBufferString("<api-error-response></api-error-response>")),
				StatusCode: http.StatusUnprocessableEntity,
				Status:     "422 Unprocessable Entity",
			},
			wantError: "422 Unprocessable Entity",
		},
		{
			name: "emptyBody",
			resp: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewBufferString("")),
				StatusCode: http.StatusUnprocessableEntity,
			},
			wantError: "EOF",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := parseError(test.resp); got == nil || got.Error() != test.wantError {
				t.Errorf("want %v, got %v", test.wantError, got)
			}
		})
	}

}
