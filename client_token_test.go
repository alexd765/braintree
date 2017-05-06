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
		_, err := bt.ClientToken().Generate(ClientTokenInput{Version: newInt(-3)})
		apiErr, ok := err.(*APIError)
		if !ok {
			t.Errorf("expected error of type APIError")
		}
		if apiErr == nil || apiErr.Code != 92806 {
			t.Errorf("api error code: got %v, want 92806", apiErr)
		}
	})
}

func newInt(n int) *int {
	return &n
}
