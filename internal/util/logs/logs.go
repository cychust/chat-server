package logs

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var loggerIns *logger

type logger struct {
	loggerMap *sync.Map
	logPath   string
	env       string // 环境变量
}

type Config struct {
	Env  string // 环境变量
	Path string // 日志文件路径
}

// 初始化
func InitLogger(config Config) error {
	if config.Path == "" {
		return errors.New("config path require")
	}
	if config.Env == "" {
		return errors.New("config env require")
	}

	loggerIns = &logger{
		loggerMap: new(sync.Map),
		env:       config.Env,
		logPath:   config.Path,
	}
	return nil
}

// 获取实例
func GetLogger(arg ...string) *logrus.Logger {
	logName := "all"
	if len(arg) > 0 {
		logName = arg[0]
	}

	if log, ok := loggerIns.loggerMap.Load(logName); ok {
		return log.(*logrus.Logger)
	}

	newLog, _ := newLogger(logName)
	loggerIns.loggerMap.Store(logName, newLog)
	return newLog
}

func newLogger(logName string) (*logrus.Logger, error) {
	if logName == "" {
		return nil, errors.New("logName require")
	}
	var logger *logrus.Logger
	if _, err := os.Stat(loggerIns.logPath); os.IsNotExist(err) {
		if err := os.MkdirAll(loggerIns.logPath, os.ModePerm); err != nil {
			return nil, err
		}
	}

	hook := lfshook.NewHook(lfshook.PathMap{
		logrus.InfoLevel:  filepath.Join(loggerIns.logPath, logName) + ".out.log",
		logrus.ErrorLevel: filepath.Join(loggerIns.logPath, logName) + ".err.log",
		logrus.DebugLevel: filepath.Join(loggerIns.logPath, logName) + ".debug.log",
		logrus.TraceLevel: filepath.Join(loggerIns.logPath, logName) + ".trace.log",
	}, &logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})
	logger = logrus.New()

	if loggerIns.env == "online" {
		logger.Out = ioutil.Discard
		logger.SetLevel(logrus.InfoLevel)
	} else {
		logger.Out = ioutil.Discard
		logger.SetLevel(logrus.DebugLevel)
	}

	logger.Hooks.Add(hook)
	return logger, nil
}
