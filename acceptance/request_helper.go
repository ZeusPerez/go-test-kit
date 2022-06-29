package acceptance

import (
	"io"
	"net/http"
	"strconv"
	"strings"
	"testing"
)

func makeRequest(t *testing.T, method, url string) (int, int) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		t.Fatalf("Failed to create request: %s", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Failed to make the request: %s", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to parse the response: %s", err)
	}

	response, err := strconv.Atoi(strings.TrimSpace(string(string(body))))
	if err != nil {
		t.Fatalf("Failed to parse the response: %s", err)
	}

	return response, resp.StatusCode
}
