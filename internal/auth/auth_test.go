package auth

import (
	"github.com/google/go-cmp/cmp"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	h := make(http.Header)
	h.Set("Authorization", "ApiKey none")
	tests := map[string]struct {
		input http.Header
		want  string
	}{
		"fail": {input: h, want: "none"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := GetAPIKey(tc.input)
			diff := cmp.Diff(tc.want, got)
			if err != nil {
				t.Fatal(err)
			}
			if diff != "" {
				t.Fatal(diff)
			}
		})
	}

}
