package mongo

import (
	"time"
	"webhook-consumer/internal/logger"
)

type WebhookError struct {
	Hm      string      `bson:"hm"`       // HTTP method
	URL     string      `bson:"url"`      // Request URL
	Hdr     interface{} `bson:"hdr"`      // Headers (Object in JS)
	Uid     int         `bson:"u_id"`     // User ID
	Swid    int         `bson:"s_wid"`    // Some webhook ID
	WType   int         `bson:"w_type"`   // Webhook type
	CidNum  string      `bson:"cid_num"`  // Caller ID number
	CallNum string      `bson:"call_num"` // Call number
	Re      int         `bson:"re"`       // Retry count
	Ac      string      `bson:"ac"`       // Action code
	CID     string      `bson:"c_id"`     // Call ID
	SipD    int         `bson:"sip_d"`    // SIP duration
	Ct      string      `bson:"ct"`       // Content type
	Pyld    interface{} `bson:"pyld"`     // Payload (Object in JS)
	Ca      time.Time   `bson:"ca"`       // Created at (Date)
	Rsp     string      `bson:"rsp"`      // Response message
	RspC    int         `bson:"rsp_c"`    // Response code
	RspT    int         `bson:"rsp_t"`    // Response time
	Rc      int         `bson:"rc"`       // Retry count or response code
	Ch      string      `bson:"ch"`       // Channel
	Offset  int         `bson:"offset"`   // Offset value
	RAt     time.Time   `bson:"r_at"`     // Received at (Date)
}

var webhook_error_log_collection = database.Collection("webhook_error_logs")

func InsertWebhookErrorLog(webhookError *WebhookError) {
	_, error := webhook_error_log_collection.InsertOne(ctx, webhookError)

	if error != nil {
		logger.Logger.Error("Can not insert webhook error log", error, webhookError)
		return
	}

	logger.Logger.Info("webhook error log saved to db")
}
