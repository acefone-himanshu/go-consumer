package consumer

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"

	kafka "github.com/segmentio/kafka-go"
)

func ProcessMessage(m kafka.Message, wg *sync.WaitGroup, sem chan struct{}) {
	defer wg.Done()
	defer func() { <-sem }()

	msg := KafkaMessage{}
	var err error = nil

	if err = json.Unmarshal(m.Value, &msg); err != nil {
		return
	}

	var payload Pyld

	switch v := msg.Pyld.(type) {
	case string:
		// Optionally, you can further unmarshal if it's JSON in a string form
		if err = json.Unmarshal([]byte(v), &payload); err != nil {
			return
		}
	case map[string]interface{}:
		// Convert the map into a structured object if needed
		jsonBytes, _ := json.Marshal(v)
		if err = json.Unmarshal(jsonBytes, &payload); err != nil {
			return
		}
	default:
		fmt.Println("Unknown pyld type")
		return
	}

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
		query.Add("test", "test")
		parsedURL.RawQuery = query.Encode()
	}

	req.URL = parsedURL

	makeApiCall(req)

	// Print the response status and body
	// fmt.Println("Response Status:", resp.Status)
	// fmt.Println("Response Body:", string(body))

}
