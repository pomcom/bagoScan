package utils

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
				"timeKey": "time"
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
