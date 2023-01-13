package utils

/*
Use those function for logging your tool
	- start
	- finish
	- warning %TODO
*/
import (
	"go.uber.org/zap"
)

func ToolStartLog(tool string, target string) {
	Logger.Info("Running", zap.String("tool", tool), zap.String("on target", target))
}

func ToolFinishedLog(tool string, target string) {
	Logger.Info("Finished", zap.String("tool", tool), zap.String("on target", target))
}

func ToolFailed(tool string, target string, err error) {
	Logger.Error("Executing failed:", zap.String("tool", tool), zap.String("on target", target), zap.Error(err))
}
