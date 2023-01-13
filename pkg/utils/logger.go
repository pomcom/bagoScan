package utils

/*
ISO8601  -> "time":1673637527.6814306"
RFC3339  -> "time":"2023-01-13T20:16:27+01:00"

*/
import (
	"encoding/json"
	"log"

	"go.uber.org/zap"
)

var Logger *zap.Logger

var rawJSON = []byte(`{
    "level": "debug",
    "encoding": "json",
    "outputPaths": ["stdout"],
    "errorOutputPaths": ["stderr"],
    "encoderConfig": {
        "messageKey": "message",
        "levelKey": "level",
        "levelEncoder": "lowercase",
				"timeKey": "time",
				"timeEncoder": "RFC3339"
    }
}`)

func InitializeLogger() {
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
