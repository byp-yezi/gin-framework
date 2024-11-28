package initialize

import (
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitZapLog() {
	switch GlobalConfig.Server.Mode {
	case "debug":
		ZapLog(zap.DebugLevel, GlobalConfig.Log.LogformatConsole)
	case "release":
		ZapLog(zap.InfoLevel, GlobalConfig.Log.LogformatConsole)
	default:
		ZapLog(zap.InfoLevel, GlobalConfig.Log.LogformatConsole)
	}
}

func ZapLog(logLevel zapcore.Level, logFormat string) {
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     GlobalConfig.Log.MessageKey,
		LevelKey:       GlobalConfig.Log.LevelKey,
		TimeKey:        GlobalConfig.Log.TimeKey,
		NameKey:        GlobalConfig.Log.NameKey,
		CallerKey:      GlobalConfig.Log.CallerKey,
		StacktraceKey:  GlobalConfig.Log.StacktraceKey,
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
	case GlobalConfig.Log.LogformatJson:
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	case GlobalConfig.Log.LogformatConsole:
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	default:
		encoder = zapcore.NewConsoleEncoder(encoderConfig)

	}

	// 添加日志切割归档功能
	hook := lumberjack.Logger{
		Filename:   GlobalConfig.Log.FileName,
		MaxSize:    GlobalConfig.Log.MaxSize,
		MaxAge:     GlobalConfig.Log.MaxAge,
		MaxBackups: GlobalConfig.Log.MaxBackups,
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
