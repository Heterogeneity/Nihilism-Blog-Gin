package core

import (
	"log"
	"os"
	"server/global"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// InitLogger 初始化并返回一个基于配置设置的新 zap.Logger 实例
func InitLogger() *zap.Logger {
	zapCfg := global.Config.Zap

	// 创建一个用于日志输出的 writeSyncer
	writeSyncer := getLogWriter(zapCfg.Filename, zapCfg.MaxSize, zapCfg.MaxBackups, zapCfg.MaxAge)

	// 如果配置了控制台输出，则添加控制台输出
	if zapCfg.IsConsolePrint {
		writeSyncer = zapcore.NewMultiWriteSyncer(writeSyncer, zapcore.AddSync(os.Stdout))
	}

	// 创建日志格式化的编码器
	encoder := getEncoder("console")

	// 根据配置确定日志级别
	var logLevel zapcore.Level

	err := logLevel.UnmarshalText([]byte(zapCfg.Level))
	if err != nil {
		log.Fatalf("解析日志等级失败: %v", err)
	}
	// 创建核心和日志实例
	core := zapcore.NewCore(encoder, writeSyncer, logLevel)
	logger := zap.New(core, zap.AddCaller())
	return logger
}

// getLogWriter 返回一个 zapcore.WriteSyncer，该写入器利用 lumberjack 包，实现日志的滚动记录
func getLogWriter(filename string, maxSize, maxBackups, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,   // 日志文件的位置
		MaxSize:    maxSize,    // 在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: maxBackups, // 保留旧文件的最大个数
		MaxAge:     maxAge,     // 保留旧文件的最大天数
	}
	return zapcore.AddSync(lumberJackLogger)
}

// getEncoder 返回一个为生产日志配置的 JSON 编码器
func getEncoder(format string) zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	encoderConfig.TimeKey = "时间"
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	//NewJSONEncoder输出格式{"level":"\u001b[34mINFO\u001b[0m","时间":"2026-04-03 16:57:11","caller":"core/server.go:24","msg":"后端已成功启动","address":"0.0.0.0:8600"}
	//NewConsoleEncoder输出格式2026-04-03 16:58:01	INFO	initialize/cron.go:17	schedule	{"keysAndValues": ["now","2026-04-03T16:58:01.7126738+08:00","entry",3,"next","2026-04-04T00:00:00+08:00"]}
	if format == "file" {
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
		return zapcore.NewJSONEncoder(encoderConfig)
	}
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}
