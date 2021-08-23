package shortener

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortLinkGenerator(t *testing.T) {
	tests := []struct {
		initialLink string
		expectshort string
		err         error
	}{
		{
			initialLink: "https://golang.org/",
			expectshort: "JJTv5dwt",
			err:         nil,
		},
		{
			initialLink: "https://www.youtube.com/watch?v=j0z4FweCy4M",
			expectshort: "BXsYD9Xw",
			err:         nil,
		},
		{
			initialLink: "https://golang.org/ | https://test.org",
			expectshort: "",
			err:         fmt.Errorf("invalid site"),
		},
	}

	testUserId := "637d8234-24a5-404d-8cd1-5968bdf38bf6"

	for _, test := range tests {
		shortlink, err := GenerateShortLink(test.initialLink, testUserId)
		assert.Equal(t, shortlink, test.expectshort)
		assert.Equal(t, err, test.err)
	}
}
