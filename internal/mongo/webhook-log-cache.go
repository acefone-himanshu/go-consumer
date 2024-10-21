package mongo

import (
	"time"
	"webhook-consumer/internal/logger"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type WebhookLogCache struct {
	CidCallNum string    `bson:"cid_call_num"` // Caller ID and Call number
	Uid        int       `bson:"u_id"`         // User ID
	Ca         time.Time `bson:"ca"`           // Created at (Date)
}

func InsertWEbhookLogCache(webhookLog *WebhookLogCache) {
	collection := client.Database(database).Collection("unique_webhook_logs_cache")

	_, error := collection.InsertOne(ctx, webhookLog)

	if error != nil {
		logger.Logger.Error("Can not insert webhook log cache", error, webhookLog)
		return
	}
}

func IsUniqueCallAttempt(cid_num string, call_num string, u_id int64) (bool, error) {
	collection := client.Database(database).Collection("unique_webhook_logs_cache")

	filter := bson.M{
		"cid_call_num": cid_num + "_" + call_num,
		"u_id":         u_id,
	}

	err := collection.FindOne(ctx, filter).Err()

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return true, nil
		}

		logger.Logger.Error("Can not find webhook log cache", err)

		return true, err
	}

	return false, nil
}
