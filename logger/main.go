package logger

import (
	"ace/model"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var lg *zap.Logger

func NewLogger(mode, name string, conf model.Logger) {
	write := setWrite(conf.Path)

	level := setLevel(conf.Level)

	encoder := setEncoder()

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)

	var core zapcore.Core
	if mode == "debug" {
		// 进入开发模式，日志输出到终端
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewTee(
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
		)
	} else {
		core = zapcore.NewCore(encoder, write, level)
	}

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段,如：添加一个服务器名称
	filed := zap.Fields(zap.String("service", name))
	// 构造日志
	lg = zap.New(core, caller, development, filed)

	zap.ReplaceGlobals(lg)
	zap.L().Info("[Service] Start logger")
}

func setWrite(path string) zapcore.WriteSyncer {
	// 日志分割
	hook := lumberjack.Logger{
		Filename:   path, // 日志文件路径，默认 os.TempDir()
		MaxSize:    10,   // 每个日志文件保存10M，默认 100M
		MaxBackups: 30,   // 保留30个备份，默认不限
		MaxAge:     7,    // 保留7天，默认不限
		Compress:   true, // 是否压缩，默认不压缩
	}
	write := zapcore.AddSync(&hook)

	return write
}

func setLevel(l string) zapcore.Level {
	// 设置日志级别
	// debug 可以打印出 info debug warn
	// info  级别可以打印 warn info
	// warn  只能打印 warn
	// debug->info->warn->error
	var level zapcore.Level
	switch l {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}

	return level
}

func setEncoder() zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "line",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	encoder := zapcore.NewJSONEncoder(encoderConfig)

	return encoder
}
