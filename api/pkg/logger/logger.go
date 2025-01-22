package logger

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type logger struct {
	logger *zap.Logger
	sugar  *zap.SugaredLogger
}

var CommonLogger *logger

func Initialize(logFormat string, logLevel zapcore.Level) error {
	CommonLogger = &logger{}

	var cfg zap.Config

	if logFormat == "json" {
		cfg = zap.Config{
			Encoding:    "json",
			Level:       zap.NewAtomicLevelAt(logLevel),
			OutputPaths: []string{"stderr"},
			EncoderConfig: zapcore.EncoderConfig{
				MessageKey:   "message",
				TimeKey:      "time",
				LevelKey:     "level",
				EncodeTime:   syslogTimeEncoder,
				EncodeLevel:  zapcore.CapitalLevelEncoder,
				EncodeCaller: zapcore.FullCallerEncoder,
			},
		}
	} else {
		cfg = zap.Config{
			Encoding:    "console",
			Level:       zap.NewAtomicLevelAt(logLevel),
			OutputPaths: []string{"stderr"},
			EncoderConfig: zapcore.EncoderConfig{
				MessageKey:   "message",
				TimeKey:      "time",
				LevelKey:     "level",
				EncodeTime:   syslogTimeEncoder,
				EncodeLevel:  customLevelEncoder,
				EncodeCaller: zapcore.FullCallerEncoder,
			},
		}
	}

	var err error
	CommonLogger.logger, err = cfg.Build()
	if err != nil {
		return err
	}

	CommonLogger.sugar = CommonLogger.logger.Sugar()
	return nil
}

func Info(msg string) {
	CommonLogger.logger.Info(msg)
}

func InfoMultiField(interfaces ...interface{}) {
	CommonLogger.sugar.Info(interfaces)
}

func Error(msg string) {
	CommonLogger.logger.Error(msg)
}

func ErrorMultiField(interfaces ...interface{}) {
	CommonLogger.sugar.Error(interfaces)
}

func Warn(msg string) {
	CommonLogger.logger.Warn(msg)
}

func Sync() error {
	if CommonLogger != nil && CommonLogger.logger != nil {
		return CommonLogger.logger.Sync()
	}
	return nil
}

func syslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

func customLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.CapitalString() + "]")
}
