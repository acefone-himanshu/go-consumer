package mongo

import (
	"time"
	"webhook-consumer/internal/logger"
)

type ErrorLog struct {
	Priority  string      `bson:"priority"`
	Message   string      `bson:"message"`
	UserID    int         `bson:"user_id"`
	Tag       string      `bson:"tag"`
	Exception interface{} `bson:"exception"`  // Use interface{} to handle dynamic objects
	Env       interface{} `bson:"env"`        // Use interface{} to handle dynamic objects
	CreatedAt time.Time   `bson:"created_at"` // Use time.Time for date fields
}

var error_log_collection = database.Collection("logs")

func InsertErrorLog(errLog *ErrorLog) {
	_, error := error_log_collection.InsertOne(ctx, errLog)

	if error != nil {
		logger.Logger.Error("Can not insert error log", error, errLog)
		return
	}

	logger.Logger.Info("error log saved to db")
}
