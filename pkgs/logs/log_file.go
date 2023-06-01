package logs

import (
	"context"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/moqsien/goutils/pkgs/gtui"
)

var (
	Logger *glog.Logger
	ctx    context.Context = context.Background()
)

func init() {
	if Logger == nil {
		Logger = glog.New()
	}
}

func SetLogger(logDir string) {
	if Logger != nil && logDir != "" {
		Logger.SetConfigWithMap(g.Map{
			"path":              logDir,
			"level":             "error",
			"stdout":            false,
			"StStatus":          0,
			"RotateSize":        "50M",
			"RotateBackupLimit": 2,
		})
	}
}

func Error(v ...interface{}) {
	if Logger != nil {
		Logger.Error(ctx, v...)
	} else {
		gtui.PrintError(v...)
	}
}

func Warning(v ...interface{}) {
	if Logger != nil {
		Logger.Warning(ctx, v...)
	} else {
		gtui.PrintWarning(v...)
	}
}

func Info(v ...interface{}) {
	if Logger != nil {
		Logger.Info(ctx, v...)
	} else {
		gtui.PrintInfo(v...)
	}
}
