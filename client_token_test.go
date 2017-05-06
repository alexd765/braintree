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
			name: "version2",
			input: ClientTokenInput{
				Version: 2,
			},
		},
		{
			name: "version3",
			input: ClientTokenInput{
				Version: 3,
			},
		},
		{
			name: "withOptions",
			input: ClientTokenInput{
				CustomerID: "cus1",
				Version:    3,
				Options: &ClientTokenOptions{
					VerifyCard: true,
				},
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
		_, err := bt.ClientToken().Generate(ClientTokenInput{})
		apiErr, ok := err.(*APIError)
		if !ok {
			t.Errorf("expected error of type APIError")
		}
		if apiErr == nil || apiErr.Code != 92806 {
			t.Errorf("api error code: got %v, want 92806", apiErr)
		}
	})
}
