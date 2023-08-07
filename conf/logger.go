package conf

import (
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitLogger() *zap.SugaredLogger {
	logMode := zapcore.DebugLevel
	if !viper.GetBool("mode.develop") {
		logMode = zapcore.InfoLevel
	}
	core := zapcore.NewCore(getEncoder(), zapcore.NewMultiWriteSyncer(getWriteSync(), zapcore.AddSync(os.Stdout)), logMode)
	return zap.New(core).Sugar()
}

// 配置日志的格式
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeTime = func(t time.Time, pae zapcore.PrimitiveArrayEncoder) {
		// 1.20+之后的时间写法,不需要记住时间了
		pae.AppendString(t.Local().Format(time.StampMilli))
	}
	return zapcore.NewJSONEncoder(encoderConfig)
}

// 配置日志的输出位置
func getWriteSync() zapcore.WriteSyncer {
	//路径分隔符,工作路径
	separator := string(filepath.Separator)
	rootDir, _ := os.Getwd()
	logFilePath := rootDir + separator + "log" + separator + time.Now().Format(time.DateOnly) + ".txt"
	lumberjackSyncer := &lumberjack.Logger{
		Filename:   logFilePath,
		MaxSize:    viper.GetInt("log.MaxSize"), // megabytes
		MaxBackups: viper.GetInt("log.MaxBackups"),
		MaxAge:     viper.GetInt("log.MaxAge"), //days
		Compress:   true,                       // disabled by default
	}
	return zapcore.AddSync(lumberjackSyncer)
}
