package httpclient

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

func Do(method, url string, headers map[string]string, body string) error {
	var reader io.Reader
	if body != "" {
		reader = bytes.NewBufferString(body)
	}

	req, err := http.NewRequest(method, url, reader)
	if err != nil {
		return err
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
		return err
	}
	defer res.Body.Close()

	elapsed := time.Since(start)
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	fmt.Println("*********** REQUEST ***********")
	fmt.Printf("%s %s\n", method, url)

	if body != "" {
		fmt.Println("Body:")
		fmt.Println(body)
	}

	fmt.Println()

	// RESPONSE
	fmt.Println("*********** RESPONSE ***********")
	fmt.Printf("STATUS %s (%dms)\n\n", res.Status, elapsed.Milliseconds())

	fmt.Println("Body:")
	fmt.Println(string(b))

	return nil
}
