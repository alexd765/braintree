package braintree

import "testing"
import "net/http"
import "bytes"
import "io/ioutil"

func TestGatewayError(t *testing.T) {
	t.Parallel()

	err := &GatewayError{Message: "Johnny"}
	got := err.Error()
	want := "Johnny"
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestValidationError(t *testing.T) {
	t.Parallel()

	err := &ValidationError{Code: 5, Message: "Johnny"}
	got := err.Error()
	want := "Code 5: Johnny"
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestProcessError(t *testing.T) {
	t.Parallel()

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
		{
			name: "invalidProcessorResponse",
			resp: &http.Response{
				Body:       ioutil.NopCloser(bytes.NewBufferString("<api-error-response><transaction><processor-response-code>text</processor-response-code></transaction></api-error-response>")),
				StatusCode: http.StatusUnprocessableEntity,
			},
			wantError: "strconv.Atoi: parsing \"text\": invalid syntax",
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
