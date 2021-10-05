package shortener

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortLinkGenerator(t *testing.T) {
	tests := []struct {
		name        string
		initialLink string
		expectshort string
		err         error
	}{
		{
			name:        "pass1",
			initialLink: "https://golang.org/",
			expectshort: "JJTv5dwt",
			err:         nil,
		},
		{
			name:        "pass2",
			initialLink: "https://www.youtube.com/watch?v=j0z4FweCy4M",
			expectshort: "BXsYD9Xw",
			err:         nil,
		},
		{
			name:        "fail",
			initialLink: "https://golang.org/ | https://test.org",
			expectshort: "",
			err:         fmt.Errorf("invalid site"),
		},
	}

	testUserId := "637d8234-24a5-404d-8cd1-5968bdf38bf6"

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			shortlink, err := GenerateShortLink(test.initialLink, testUserId)
			assert.Equal(t, shortlink, test.expectshort)
			assert.Equal(t, err, test.err)
		})
	}
}
