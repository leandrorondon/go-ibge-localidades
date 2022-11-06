package ibgelocalidades_test

import (
	"net/http"
	"testing"
	"time"

	ibgelocalidades "github.com/leandrorondon/go-ibge-localidades"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	to := time.Second
	tests := []struct {
		name      string
		roundTrip http.RoundTripper
		timeout   *time.Duration
	}{
		{
			name: "no options",
		},
		{
			name:      "with transport",
			roundTrip: &http.Transport{},
		},
		{
			name:    "with timeout",
			timeout: &to,
		},
		{
			name:      "with timeout and transport",
			timeout:   &to,
			roundTrip: &http.Transport{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var options []ibgelocalidades.Option
			if test.roundTrip != nil {
				options = append(options, ibgelocalidades.WithTransport(test.roundTrip))
			}
			if test.timeout != nil {
				options = append(options, ibgelocalidades.WithTimeout(*test.timeout))
			}

			client := ibgelocalidades.New(options...)
			assert.NotNil(t, client)
		})
	}
}
