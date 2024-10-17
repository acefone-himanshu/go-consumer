package mongo

import (
	"time"
	"webhook-consumer/internal/logger"
)

type ProviderErrorLog struct {
	Priority  string      `bson:"priority"`
	Message   string      `bson:"message"`
	UserID    int         `bson:"user_id"`
	Tag       string      `bson:"tag"`
	Exception interface{} `bson:"exception"`
	Env       interface{} `bson:"env"`
	CreatedAt time.Time   `bson:"created_at"`
}

var provider_error_log_collection = database.Collection("provider_error_logs")

func InsertProviderErrorLog(errLog *ProviderErrorLog) {
	_, error := provider_error_log_collection.InsertOne(ctx, errLog)

	if error != nil {
		logger.Logger.Error("Can not insert provider error log", error, errLog)
		return
	}

	logger.Logger.Info("provider error log saved to db")
}
