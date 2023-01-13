package utils

/*
Only part where the native go logger (for example: `log.Fatalf()` used should be in this part.
Use `utils.Logger` everwhere else:

 utils.Logger.Info("Some Information")
 utils.Logger.Warn("Warning")
 utils.Logger.Error("Error")

ISO8601  -> "time":1673637527.6814306"
RFC3339  -> "time":"2023-01-13T20:16:27+01:00"


*/
import (
	"encoding/json"
	"log"
	"os"

	"go.uber.org/zap"
)

var Logger *zap.Logger
var logDir = "/output/logs"

var rawJSON = []byte(`{
    "level": "debug",
    "encoding": "json",
		"outputPaths": ["stdout", "` + logDir + `/bargoScan.log"],
    "errorOutputPaths": ["stderr", "` + logDir + `/bargoScan.log"],
    "encoderConfig": {
        "messageKey": "message",
        "levelKey": "level",
        "levelEncoder": "lowercase",
				"timeKey": "time",
				"timeEncoder": "RFC3339"
    }
}`)

func InitializeLogger() {

	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		os.MkdirAll(logDir, os.ModePerm)
		log.Println("Output directory has been created")
	}

	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		log.Fatalf("Failed to parse zap config: %v", err)
	}
	var err error
	Logger, err = cfg.Build()
	if err != nil {
		log.Fatalf("Failed to build zap logger: %v", err)
	}
	defer Logger.Sync()
}
