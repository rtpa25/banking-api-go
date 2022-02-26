package logger

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	//load the env variables
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	//custom logger implemantation
	var err error
	config := zap.NewProductionConfig()
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time-stamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.StacktraceKey = ""
	config.EncoderConfig = encoderConfig
	log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err.Error())
	}
}

func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	log.Error(message, fields...)
}
