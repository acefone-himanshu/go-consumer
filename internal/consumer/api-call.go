package consumer

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const timeout = time.Second * 10

func prepareRequest(msg *KafkaMessage, payload Pyld) {
	var err error
	var requestBody []byte

	if requestBody, err = json.Marshal(payload); err != nil {
		return
	}

	var req *http.Request
	if req, err = http.NewRequest(strings.ToUpper(msg.Hm), msg.URL, nil); err != nil {
		return
	}

	var parsedURL *url.URL

	if parsedURL, err = url.Parse(req.URL.String()); err != nil {
		return
	}

	// loop through headers
	for k, v := range msg.Hdr {
		req.Header.Add(k, fmt.Sprint(v))
	}

	if msg.Hm == "POST" {
		if msg.Hdr["content-type"] == "application/json" {
			req.Body = io.NopCloser(strings.NewReader(string(requestBody)))
		} else if msg.Hdr["content-type"] == "application/x-www-form-urlencoded" {
			values := url.Values{}
			for k, v := range payload {
				values.Add(k, fmt.Sprint(v))
			}
			req.Body = io.NopCloser(strings.NewReader(values.Encode()))
		}
	} else {
		query := parsedURL.Query()
		for k, v := range payload {
			query.Add(k, fmt.Sprint(v))
		}
		parsedURL.RawQuery = query.Encode()
	}

	req.URL = parsedURL

	var res *http.Response

	start_time := time.Now()
	res, err = makeApiCall(req)

	elapsed := time.Since(start_time)

	if err != nil {

	}

	var body []byte

	if body, err = io.ReadAll(res.Body); err != nil {
	}

	msg.rsp_c = res.StatusCode
	msg.rsp_t = int(elapsed.Seconds())
	msg.rsp = string(body)

	_ = res
	// Print the response status and body
	// fmt.Println("Response Status:", resp.Status)
	// fmt.Println("Response Body:", string(body))

}

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

	// if body, err = io.ReadAll(resp.Body); err != nil {
	// 	return resp, err
	// }

	_ = body

	return resp, err
}
