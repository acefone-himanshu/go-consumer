package mongo

import (
	"time"
	"webhook-consumer/internal/logger"
)

type Webhook struct {
	Hm      string      `bson:"hm"`       // HTTP method
	URL     string      `bson:"url"`      // Request URL
	Hdr     interface{} `bson:"hdr"`      // Headers (Object in JS)
	Uid     int64       `bson:"u_id"`     // User ID
	Swid    int64       `bson:"s_wid"`    // Some webhook ID
	WType   int         `bson:"w_type"`   // Webhook type
	CidNum  string      `bson:"cid_num"`  // Caller ID number
	CallNum string      `bson:"call_num"` // Call number
	Re      int64       `bson:"re"`       // Retry count
	Ac      string      `bson:"ac"`       // Action code
	CID     string      `bson:"c_id"`     // Call ID
	SipD    uint32      `bson:"sip_d"`    // SIP duration
	Ct      time.Time   `bson:"ct"`       // Call time
	Pyld    interface{} `bson:"pyld"`     // Payload (Object in JS)
	Ca      time.Time   `bson:"ca"`       // Created at (Date)
	Rsp     string      `bson:"rsp"`      // Response message
	RspC    int         `bson:"rsp_c"`    // Response code
	RspT    float64     `bson:"rsp_t"`    // Response time
	Rc      int         `bson:"rc"`       // Retry count or response code
	Ch      string      `bson:"ch"`       // Channel
	Offset  int64       `bson:"offset"`   // Offset value
	Meta    interface{} `bson:"meta"`     // Meta (Object in JS)
	E       interface{} `bson:"e"`        // Mixed type (Schema.Types.Mixed)
}

func InsertWebhookLog(webhook *Webhook) {
	collection := client.Database(database).Collection("webhook_logs_new_go")

	_, error := collection.InsertOne(ctx, webhook)

	if error != nil {
		logger.Logger.Error("Can not insert webhook", error, webhook)
		return
	}
}
