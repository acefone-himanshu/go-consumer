package consumer

import (
	"net/http"
	"time"
)

const timeout = time.Second * 10

func makeApiCall(req *http.Request) (*http.Response, error) {
	client := &http.Client{
		Timeout: timeout,
	}

	var err error
	var resp *http.Response

	if resp, err = client.Do(req); err != nil {
		return resp, err
	}

	defer resp.Body.Close()

	var body []byte

	// if body, err = io.ReadAll(resp.Body); err != nil {
	// 	return resp, err
	// }

	_ = body

	return resp, err
}
