package initialize

import (
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"gin-framework/config"
)

func InitZapLog() {
	switch config.GlobalConfig.Server.Mode {
	case "debug":
		ZapLog(zap.DebugLevel, config.GlobalConfig.Log.LogformatConsole)
	case "release":
		ZapLog(zap.InfoLevel, config.GlobalConfig.Log.LogformatConsole)
	default:
		ZapLog(zap.InfoLevel, config.GlobalConfig.Log.LogformatConsole)
	}
}

func ZapLog(logLevel zapcore.Level, logFormat string) {
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     config.GlobalConfig.Log.MessageKey,
		LevelKey:       config.GlobalConfig.Log.LevelKey,
		TimeKey:        config.GlobalConfig.Log.TimeKey,
		NameKey:        config.GlobalConfig.Log.NameKey,
		CallerKey:      config.GlobalConfig.Log.CallerKey,
		StacktraceKey:  config.GlobalConfig.Log.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.MillisDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 设置日志输出格式
	var encoder zapcore.Encoder
	switch logFormat {
	case config.GlobalConfig.Log.LogformatJson:
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	case config.GlobalConfig.Log.LogformatConsole:
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	default:
		encoder = zapcore.NewConsoleEncoder(encoderConfig)

	}

	// 添加日志切割归档功能
	hook := lumberjack.Logger{
		Filename:   config.GlobalConfig.Log.FileName,
		MaxSize:    config.GlobalConfig.Log.MaxSize,
		MaxAge:     config.GlobalConfig.Log.MaxAge,
		MaxBackups: config.GlobalConfig.Log.MaxBackups,
		Compress:   true,
	}

	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stderr), zapcore.AddSync(&hook)),
		zap.NewAtomicLevelAt(logLevel),
	)

	caller := zap.AddCaller()
	development := zap.Development()
	logger := zap.New(core, caller, development)
	zap.ReplaceGlobals(logger)
}
