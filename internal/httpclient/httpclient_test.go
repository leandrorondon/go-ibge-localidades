package httpclient_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/leandrorondon/go-ibge-localidades/internal/httpclient"
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
			var options []httpclient.ClientOption
			if test.roundTrip != nil {
				options = append(options, httpclient.WithTransport(test.roundTrip))
			}
			if test.timeout != nil {
				options = append(options, httpclient.WithTimeout(*test.timeout))
			}

			client := httpclient.New(options...)
			assert.NotNil(t, client)
		})
	}
}

func TestGet(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK")
	}))
	defer srv.Close()

	tests := []struct {
		name             string
		url              string
		expectedResponse []byte
		expectError      bool
	}{
		{
			name:        "invalid url",
			url:         string([]byte{0x7f}),
			expectError: true,
		},
		{
			name:        "invalid protocol",
			url:         "inv://protocol",
			expectError: true,
		},
		{
			name:             "success",
			url:              srv.URL,
			expectError:      false,
			expectedResponse: []byte("OK"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			resp, err := httpclient.New().Get(context.Background(), test.url)
			if test.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, test.expectedResponse, resp)
		})
	}
}

func TestPost(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK")
	}))
	defer srv.Close()

	srvErrResponse := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Length", "1")
	}))
	defer srv.Close()

	tests := []struct {
		name             string
		url              string
		body             any
		expectedResponse []byte
		expectError      bool
	}{
		{
			name:        "invalid url",
			url:         string([]byte{0x7f}),
			expectError: true,
		},
		{
			name:        "invalid protocol",
			url:         "inv://protocol",
			expectError: true,
		},
		{
			name:        "invalid body",
			url:         srv.URL,
			body:        make(chan int),
			expectError: true,
		},
		{
			name:        "error reading response",
			url:         srvErrResponse.URL,
			expectError: true,
		},
		{
			name:             "success",
			url:              srv.URL,
			expectError:      false,
			expectedResponse: []byte("OK"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			resp, err := httpclient.New().Post(context.Background(), test.url, test.body)
			if test.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, test.expectedResponse, resp)
		})
	}
}
