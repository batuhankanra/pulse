package httpclient

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func Do(method, url string, headers map[string]string, body string) (string, error) {
	var reader io.Reader
	if body != "" {
		reader = bytes.NewBufferString(body)
	}
	req, err := http.NewRequest(method, url, reader)
	if err != nil {
		return "", err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	client := &http.Client{
		Timeout: 15 * time.Second,
	}
	start := time.Now()
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	elapsed := time.Since(start)
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	var out strings.Builder

	out.WriteString(fmt.Sprintf("STATUS %d %s (%dms)\n\n",
		res.StatusCode,
		res.Status,
		elapsed.Milliseconds()))
	out.WriteString("Headers:\n")
	for k, v := range res.Header {
		out.WriteString(fmt.Sprintf("  %s: %s\n", k, strings.Join(v, ",")))

	}
	out.WriteString("\nBody:\n")
	out.Write(b)

	return out.String(), nil

}
