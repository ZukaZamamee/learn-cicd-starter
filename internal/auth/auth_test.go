package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	cases := []struct {
		input          http.Header
		expectedAPIKey string
		expectedError  error
	}{
		{
			input:          http.Header{"Authorization": []string{"ApiKey Bearer-token"}},
			expectedAPIKey: "Bearer-token",
			expectedError:  nil,
		},
		{
			input:          http.Header{"Authorization": []string{"AuthKey Bearer-token"}},
			expectedAPIKey: "",
			expectedError:  errors.New("malformed authorization header"),
		},
		{
			input:          http.Header{"Auth": []string{"AuthKey Bearer-token"}},
			expectedAPIKey: "",
			expectedError:  errors.New("no authorization header included"),
		},
	}

	for _, c := range cases {
		ApiKey, err := GetAPIKey(c.input)
		if ApiKey != c.expectedAPIKey {
			t.Errorf("expectedAPIKey doesn't match: '%v' vs '%v'", ApiKey, c.expectedAPIKey)
		}
		if c.expectedError == nil && err != nil {
			t.Errorf("expected no error but got error: '%v' vs expected '%v'", err.Error(), c.expectedError)
		}
		if c.expectedError != nil && err == nil {
			t.Errorf("expected error but got nil: expected '%v'", c.expectedError)
		}
		if c.expectedError != nil && err != nil && err.Error() != c.expectedError.Error() {
			t.Errorf("expectedError doesn't match: '%v' vs '%v'", err.Error(), c.expectedError.Error())
		}
	}

}
