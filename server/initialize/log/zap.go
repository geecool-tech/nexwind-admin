package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

// InitZap 初始化 zap logger
func InitZap() *zap.Logger {
	colorConfig := zap.NewDevelopmentConfig()
	colorConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	colorConfig.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05.999"))
	}
	consoleEncoder := zapcore.NewConsoleEncoder(colorConfig.EncoderConfig)
	consoleCore := zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), colorConfig.Level)
	fileEncoder := zapcore.NewJSONEncoder(colorConfig.EncoderConfig)
	infoFileCore := zapcore.NewCore(fileEncoder, zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./logs/info.log",
		MaxSize:    100, // MB
		MaxBackups: 10,
		MaxAge:     30, // Days
	}), zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.InfoLevel
	}))
	warnFileCore := zapcore.NewCore(fileEncoder, zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./logs/warn.log",
		MaxSize:    100, // MB
		MaxBackups: 10,
		MaxAge:     30, // Days
	}), zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.WarnLevel
	}))
	errorFileCore := zapcore.NewCore(fileEncoder, zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./logs/error.log",
		MaxSize:    100, // MB
		MaxBackups: 10,
		MaxAge:     30, // Days
	}), zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.ErrorLevel
	}))
	logger := zap.New(zapcore.NewTee(consoleCore, infoFileCore, warnFileCore, errorFileCore))
	return logger
}
