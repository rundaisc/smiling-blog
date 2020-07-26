package slog

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"path"
	"time"
)

var Log *zap.SugaredLogger
var Slog LogConfig

type LogConfig struct {
	Level    string `yaml:"level"`
	Path     string `yaml:"path"`
	FileName string `yaml:"filename"`
}

func InitLogger(logConfig LogConfig) {
	encoder := getEncoder()
	//info level
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})
	infoWriter := getLogWriter(logConfig.Path, logConfig.FileName+"-info")
	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	//error level
	errorWriter := getLogWriter(logConfig.Path, logConfig.FileName+"-error")
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zap.DebugLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(errorWriter), errorLevel),
	)
	// develop mode
	caller := zap.AddCaller()
	// open the code line
	development := zap.Development()
	logger := zap.New(core, caller, development, zap.AddCallerSkip(1))
	Log = logger.Sugar()
}

/**
 * time format
 */
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("[2006-01-02 15:04:05]"))
}

/**
 * get zap log encoder
 */
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = customTimeEncoder
	encoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
	encoderConfig.LineEnding = zapcore.DefaultLineEnding
	return zapcore.NewConsoleEncoder(encoderConfig)
}
func getLogWriter(logPath, level string) io.Writer {
	logFullPath := path.Join(logPath, level)
	hook, err := rotatelogs.New(
		logFullPath+"-%Y%m%d%H"+".txt",
		// log file split
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		panic(err)
	}
	return hook
}
func (slf *LogConfig) Print(v ...interface{}) {
	Log.Info(v)
}
func Info(args ...interface{}) {
	Log.Info(args)
}
func Infof(template string, args ...interface{}) {
	Log.Infof(template, args...)
}
func Debug(args ...interface{}) {
	Log.Debug(args)
}
func Debugf(template string, args ...interface{}) {
	Log.Debugf(template, args...)
}
func Error(args ...interface{}) {
	Log.Error(args)
}
func Errorf(template string, args ...interface{}) {
	Log.Errorf(template, args...)
}
