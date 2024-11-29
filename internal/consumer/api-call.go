package consumer

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
	"webhook-consumer/internal/mongo"
)

func prepareRequest(msg *KafkaMessage, payload *Pyld) {
	var err error
	var requestBody []byte

	disableTimeoutUserIdStr := os.Getenv("DISABLE_TIMEOUT_USER_ID")
	disableTimeoutUserId := ""

	if msg.UID != 0 {
		disableTimeoutUserId = strconv.FormatInt(msg.UID, 10)
	}

	timeout := time.Second * 10

	if disableTimeoutUserIdStr == disableTimeoutUserId {
		timeout = 0
	}

	if requestBody, err = json.Marshal(payload); err != nil {
		return
	}

	var req *http.Request
	if req, err = http.NewRequest(strings.ToUpper(msg.Hm), msg.URL, nil); err != nil {
		return
	}

	var parsedURL *url.URL

	if parsedURL, err = url.Parse(msg.URL); err != nil {
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
			for k, v := range *payload {
				values.Add(k, fmt.Sprint(v))
			}
			req.Body = io.NopCloser(strings.NewReader(values.Encode()))
		}
	} else {
		query := parsedURL.Query()
		for k, v := range *payload {
			query.Add(k, fmt.Sprint(v))
		}
		parsedURL.RawQuery = query.Encode()
	}

	req.URL = parsedURL

	start_time := time.Now()

	// make api call
	client := &http.Client{
		Timeout: timeout,
	}

	var resp *http.Response

	if resp, err = client.Do(req); err != nil || true {
		return
	}

	defer resp.Body.Close()

	elapsed := time.Since(start_time)

	var body []byte

	if body, err = io.ReadAll(resp.Body); err != nil {
	}

	msg.rsp_c = resp.StatusCode
	msg.rsp_t = elapsed.Seconds()
	msg.rsp = string(body)

	if resp.StatusCode >= 400 && msg.Re == 1 && msg.Rc < 5 {
		saveLogForRetries(msg)
	}

	webhookLog := &mongo.Webhook{
		Hm:      msg.Hm,
		URL:     msg.URL,
		Hdr:     msg.Hdr,
		Uid:     msg.UID,
		Swid:    msg.SWid,
		WType:   msg.WType,
		CidNum:  msg.CID,
		CallNum: msg.CallNum,
		Re:      msg.Re,
		Ac:      msg.AC,
		CID:     msg.CID,
		SipD:    ipToNumber(msg.SIPD),
		Pyld:    msg.Pyld,
		Ca:      time.Now(),
		Rsp:     msg.rsp,
		RspC:    msg.rsp_c,
		RspT:    msg.rsp_t,
		Rc:      msg.rsp_c,
		Ch:      msg.Ch,
		Offset:  msg.offset,
		Meta:    msg.Meta,
		E:       err,
	}

	mongo.InsertWebhookLog(webhookLog)

}

func saveLogForRetries(msg *KafkaMessage) {
	// TODO : remove
	if true {
		return
	}

	now := time.Now()
	r_at := now.Add(time.Hour * 1)
	msg.r_at = r_at

	if msg.Rc == 0 {
		msg.Rc = 0
	} else {
		msg.Rc = msg.Rc + 1
	}

	webhookErrorLog := &mongo.WebhookError{
		Hm:      msg.Hm,
		URL:     msg.URL,
		Hdr:     msg.Hdr,
		Uid:     int(msg.UID),
		Swid:    int(msg.SWid),
		WType:   int(msg.WType),
		CidNum:  msg.CidNum,
		CallNum: msg.CallNum,
		Re:      int(msg.Re),
		Ac:      msg.AC,
		CID:     msg.CID,
		SipD:    int(ipToNumber(msg.SIPD)),
		Ct:      "",
		Pyld:    msg.Pyld,
		Ca:      msg.ca,
		Rsp:     msg.rsp,
		RspC:    msg.rsp_c,
		RspT:    msg.rsp_t,
		Rc:      int(msg.Rc),
		Ch:      msg.Ch,
		Offset:  int(msg.offset),
		RAt:     msg.r_at,
	}
	mongo.InsertWebhookErrorLog(webhookErrorLog)
}

func ipToNumber(ipString string) uint32 {
	ip := net.ParseIP(ipString)
	if ip == nil {
		return 0
	}

	// Ensure it's an IPv4 address
	ip = ip.To4()
	if ip == nil {
		return 0
	}

	// Convert to uint32
	result := uint32(ip[0])<<24 | uint32(ip[1])<<16 | uint32(ip[2])<<8 | uint32(ip[3])

	return result
}
