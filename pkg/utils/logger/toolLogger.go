package utils

/*
Use those function for logging your tool
*/
import (
	"os/exec"

	"go.uber.org/zap"
)

func ToolStartLog(tool string, target string) {
	Logger.Info("running", zap.String("tool", tool), zap.String("target", target))
}

func ToolFinishedLog(tool string, target string) {
	Logger.Info("finished", zap.String("tool", tool), zap.String("target", target))
}

func ToolFailed(tool string, target string, err error) {
	Logger.Error("execution failed", zap.String("tool", tool), zap.String("target", target), zap.Error(err))
}

func ExecutedCommand(cmd *exec.Cmd) {
	Logger.Info("full command", zap.String("path", cmd.Path), zap.Strings("args", cmd.Args[1:]))
}
