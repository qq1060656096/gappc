package generate

const BootstrapTemplate = `package bootstrap

import (
	"github.com/joho/godotenv"
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/qq1060656096/common"
	"github.com/sirupsen/logrus"
	"io"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

type Application struct {
	LogWriter   io.Writer
	dir         string
	logDir      string
	isBootstrap bool
}

var app *Application

func App() *Application {
	if app != nil {
		return app
	}
	app := NewApplication()
	app.Bootstrap()
	return app
}

func NewApplication() *Application {
	app := &Application{}
	return app
}

func (a *Application) Bootstrap() *Application {
	if a.isBootstrap {
		return a
	}
	// 设置应用目录
	a.setDir()
	// 加载配置
	godotenv.Load(a.dir+"/config/.app.env", a.dir+"/config/.db.env", a.dir+"/config/.cache.env")
	// 设置日志目录
	a.setLogDir()

	// 初始化数据库和redis连接
	common.DbConnInit()
	common.RedisConnInit()

	a.isBootstrap = true
	return a
}

func (a *Application) setDir() *Application {
	if a.isBootstrap {
		return a
	}
	_, fileName, _, ok := runtime.Caller(1)
	if !ok {
		logrus.Errorf("bootstrap.app.dir.failed: %s", ok)
	}
	dir, err := filepath.Abs(filepath.Dir(filepath.Dir(fileName)))
	if err != nil {
		logrus.Errorf("bootstrap.app.dir.abs.failed: %s", err)
		return nil
	}
	a.dir = dir
	logrus.Info("bootstrap.app.dir.abs", a.dir)
	return a
}

func (a *Application) setLogDir() *Application {
	if a.isBootstrap {
		return a
	}
	// 设置日志路径
	logDir := common.OsEnvManager.Get("LOG_DIR")
	a.logDir = strings.TrimRight(logDir, "/")
	if a.logDir == "" {
		a.logDir = a.dir + "/runtime/log"
	}
	logrus.Info("bootstrap.app.logDir.abs", a.logDir)

	// 设置分割日志
	logWriterFormat := a.logDir + "/gapp.%Y%m%d.log"
	logWriter, err := rotatelogs.New(
		logWriterFormat,
		rotatelogs.WithLinkName(a.logDir+"/gapp.access.log"),
		rotatelogs.WithMaxAge(24*time.Hour),
		rotatelogs.WithRotationTime(time.Hour),
	)
	if err != nil {
		logrus.Errorf("bootstrap.rotatelogs.New.failed: %s", err)
		return nil
	}
	logrus.SetOutput(logWriter)
	a.LogWriter = logWriter
	return a
}
`