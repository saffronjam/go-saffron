package log

import (
	"github.com/saffronjam/go-saffron/pkg/core"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var OnLog core.SubscriberList[zapcore.Entry]

var (
	Logger    *zap.SugaredLogger
	LoggerMap = make(map[string]*zap.SugaredLogger)
)

var (
	defaultLogger = "default"
)

type triggerCore struct {
	zapcore.Core
}

func (c *triggerCore) With(fields []zapcore.Field) zapcore.Core {
	return &triggerCore{
		Core: c.Core.With(fields),
	}
}

func (c *triggerCore) Check(entry zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if c.Enabled(entry.Level) {
		return ce.AddCore(entry, c)
	}
	return ce
}

func (c *triggerCore) Write(entry zapcore.Entry, fields []zapcore.Field) error {
	err := c.Core.Write(entry, fields)
	OnLog.Trigger(entry)
	return err
}

func SetupLogger() error {
	Logger = Get(defaultLogger)
	return nil
}

func Get(name string) *zap.SugaredLogger {
	if sugaredLogger, ok := LoggerMap[name]; ok {
		return sugaredLogger
	}

	var zapCore zapcore.Core
	var encoderCfg zapcore.EncoderConfig
	var level zapcore.LevelEnabler

	encoderCfg = zap.NewDevelopmentEncoderConfig()
	level = zapcore.DebugLevel

	consoleEncoder := zapcore.NewConsoleEncoder(encoderCfg)
	consoleSyncer := zapcore.AddSync(zapcore.Lock(os.Stdout))
	baseCore := zapcore.NewCore(consoleEncoder, consoleSyncer, level)

	zapCore = &triggerCore{Core: baseCore}

	logger := zap.New(zapCore, zap.WithCaller(false))
	if name == defaultLogger {
		LoggerMap[name] = logger.Sugar()
	} else {
		LoggerMap[name] = logger.Sugar().Named(name)
	}

	return LoggerMap[name]
}

func Logln(lvl zapcore.Level, args ...interface{}) {
	Logger.Logln(lvl, args...)
}

func Debugln(args ...interface{}) {
	Logger.Debugln(args...)
}

func Println(args ...interface{}) {
	Logger.Infoln(args...)
}

func Infoln(args ...interface{}) {
	Logger.Infoln(args...)
}

func Warnln(args ...interface{}) {
	Logger.Warnln(args...)
}

func Errorln(args ...interface{}) {
	Logger.Errorln(args...)
}

func Fatalln(args ...interface{}) {
	Logger.Fatalln(args...)
}

func Logf(lvl zapcore.Level, template string, args ...interface{}) {
	Logger.Logf(lvl, template, args...)
}

func Debugf(template string, args ...interface{}) {
	Logger.Debugf(template, args...)
}

func Printf(template string, args ...interface{}) {
	Logger.Infof(template, args...)
}

func Infof(template string, args ...interface{}) {
	Logger.Infof(template, args...)
}

func Warnf(template string, args ...interface{}) {
	Logger.Warnf(template, args...)
}

func Errorf(template string, args ...interface{}) {
	Logger.Errorf(template, args...)
}

func Fatalf(template string, args ...interface{}) {
	Logger.Fatalf(template, args...)
}
