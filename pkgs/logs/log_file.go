package logs

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gvcgo/goutils/pkgs/gtea/gprint"
	utils "github.com/gvcgo/goutils/pkgs/gutils"
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
	utils.MakeDirs(logDir)
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
		gprint.PrintError(fmt.Sprint(v...))
	}
}

func Warning(v ...interface{}) {
	if Logger != nil {
		Logger.Warning(ctx, v...)
	} else {
		gprint.PrintWarning(fmt.Sprint(v...))
	}
}

func Info(v ...interface{}) {
	if Logger != nil {
		Logger.Info(ctx, v...)
	} else {
		gprint.PrintInfo(fmt.Sprint(v...))
	}
}
