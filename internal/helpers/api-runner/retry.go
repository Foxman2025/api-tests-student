package api_runner

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

var retryHTTPClient = &http.Client{Transport: retryTransport{}}

type retryTransport struct{}

// RoundTrip один раз повторяет запрос при 503.
func (retryTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	var body []byte
	if req.Body != nil {
		defer req.Body.Close()

		var err error
		body, err = io.ReadAll(req.Body)
		if err != nil {
			return nil, err
		}
	}

	send := func() (*http.Response, error) {
		attempt := req.Clone(req.Context())
		attempt.Body = http.NoBody
		if len(body) > 0 {
			attempt.Body = io.NopCloser(bytes.NewReader(body))
		}
		return http.DefaultTransport.RoundTrip(attempt)
	}

	response, err := send()
	if err != nil || response.StatusCode != http.StatusServiceUnavailable {
		return response, err
	}

	const maxDelay = 10 * time.Second
	delay := 5 * time.Second
	retryAfter := response.Header.Get("Retry-After")

	if seconds, err := strconv.ParseUint(retryAfter, 10, 64); err == nil {
		if seconds > uint64(maxDelay/time.Second) {
			return response, nil
		}
		delay = max(delay, time.Duration(seconds)*time.Second)
	} else if retryAt, err := http.ParseTime(retryAfter); err == nil {
		delay = max(delay, time.Until(retryAt))
	}

	if delay > maxDelay {
		return response, nil
	}

	_ = response.Body.Close()
	log.Printf("Получили 503. Повторяем %s %s через %s", req.Method, req.URL.Path, delay)

	select {
	case <-req.Context().Done():
		return nil, req.Context().Err()
	case <-time.After(delay):
		return send()
	}
}
