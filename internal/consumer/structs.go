package consumer

type KafkaMessage struct {
	UID       int64       `json:"u_id"`
	CID       string      `json:"c_id"`
	AC        string      `json:"ac"`
	SIPD      string      `json:"sip_d"`
	SWid      int64       `json:"s_wid"`
	URL       string      `json:"url"`
	Hm        string      `json:"hm"`
	Hdr       Hdr         `json:"hdr"`
	WType     int64       `json:"w_type"`
	Re        int64       `json:"re"`
	CidNum    string      `json:"cid_num"`
	CallNum   string      `json:"call_num"`
	Pyld      interface{} `json:"pyld"`
	Meta      Meta        `json:"meta"`
	Ch        string      `json:"ch"`
	Partition int64       `json:"partition"`
}

type BillingCircle struct {
	Operator string `json:"operator"`
	Circle   string `json:"circle"`
}

type Hdr = map[string]interface{}
type Meta = map[string]interface{}
type Pyld = map[string]interface{}
