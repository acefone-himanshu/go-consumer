package mongo

import (
	"time"
	"webhook-consumer/internal/logger"
)

type ProviderAPI struct {
	HM      string      `bson:"hm"`
	URL     string      `bson:"url"`
	HDR     interface{} `bson:"hdr"` // Use interface{} for dynamic objects
	UID     int         `bson:"u_id"`
	SWID    int         `bson:"s_wid"`
	WType   int         `bson:"w_type"`
	CIDNum  string      `bson:"cid_num"`
	CallNum string      `bson:"call_num"`
	RE      int         `bson:"re"`
	AC      string      `bson:"ac"`
	CID     string      `bson:"c_id"`
	SIPD    int         `bson:"sip_d"`
	CT      string      `bson:"ct"`
	Pyld    interface{} `bson:"pyld"` // Use interface{} for dynamic objects
	CA      time.Time   `bson:"ca"`   // Use time.Time for date fields
	RSP     string      `bson:"rsp"`
	RSPC    int         `bson:"rsp_c"`
	RSPT    int         `bson:"rsp_t"`
	RC      int         `bson:"rc"`
	CH      string      `bson:"ch"`
	Offset  int         `bson:"offset"`
	PID     int         `bson:"p_id"`   // Required field
	PType   int         `bson:"p_type"` // Required field
	Meta    interface{} `bson:"meta"`   // Use interface{} for dynamic objects
	E       interface{} `bson:"e"`      // Use interface{} for dynamic objects
}

func InsertProviderWebhook(providerAPI *ProviderAPI) {
	collection := client.Database(database).Collection("provider_api")

	_, err := collection.InsertOne(ctx, providerAPI)
	if err != nil {
		logger.Logger.Error("Can not insert provider api", err, providerAPI)
		return
	}

	logger.Logger.Info("provider api saved to db")
}
