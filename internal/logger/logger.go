package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
)

var Logger = log.New()

func init() {
	// Log as JSON instead of the default ASCII formatter.
	Logger.SetFormatter(&log.TextFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	Logger.SetOutput(os.Stdout)

	// You could set this to any `io.Writer` such as a file
	// file, err := os.OpenFile("logs/logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// if err == nil {
	// 	Logger.Out = file
	// } else {
	// 	Logger.Info("Failed to log to file, using default stderr")
	// }

}
