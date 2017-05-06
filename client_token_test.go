package braintree

import "testing"

func TestGenerateClientToken(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     ClientTokenInput
		wantError error
	}{
		{
			name:  "minimal",
			input: ClientTokenInput{},
		},
		{
			name: "version2",
			input: ClientTokenInput{
				Version: newInt(2),
			},
		},
		{
			name: "version3",
			input: ClientTokenInput{
				Version: newInt(3),
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got, err := bt.ClientToken().Generate(test.input)
			if err != nil {
				t.Fatalf("unexpected err: %s", err)
			}
			if len(got) < 10 {
				t.Errorf("Client Token: got '%v', want a longer one.", got)
			}
		})
	}

	t.Run("invalidVersion", func(t *testing.T) {
		t.Parallel()
		if _, err := bt.ClientToken().Generate(ClientTokenInput{Version: newInt(-3)}); err == nil || err.Error() != "422 Unprocessable Entity" {
			t.Errorf("got: %v, want: 422 Unprocessable Entity", err)
		}
	})
}

func newInt(n int) *int {
	return &n
}
