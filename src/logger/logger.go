package logger

import (
	"errors"
	"noraclock/src/configs"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var conf = configs.Get()

var generalLogger *zap.Logger
var accessLogger *zap.Logger

var levelMap = map[string]zapcore.Level{
	"debug": zap.DebugLevel,
	"info":  zap.InfoLevel,
	"warn":  zap.WarnLevel,
	"error": zap.ErrorLevel,
}

// General returns the general logger.
func General() *zap.Logger {
	if generalLogger != nil {
		return generalLogger
	}

	var err error
	generalLogger, err = createLogger(conf.Logger.Level, conf.Logger.GeneralFilePath)
	if err != nil {
		panic(err)
	}
	return generalLogger
}

// Access returns the access logger.
func Access() *zap.Logger {
	if accessLogger != nil {
		return accessLogger
	}

	var err error
	accessLogger, err = createLogger(conf.Logger.Level, conf.Logger.AccessFilePath)
	if err != nil {
		panic(err)
	}
	return accessLogger
}

// createLogger creates a zap logger using the provided arguments.
// The created logger does not log Service Name, PID and Hostname by default.
func createLogger(logLevel string, filePath string) (*zap.Logger, error) {
	fileEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	file, err := getFilePointer(filePath)
	if err != nil {
		return nil, err
	}

	zapLogLevel, exists := levelMap[logLevel]
	if !exists {
		return nil, errors.New("unknown log level provided")
	}

	levelEnabler := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapLogLevel
	})

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, zapcore.Lock(file), levelEnabler),
		zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), levelEnabler),
	)

	return zap.New(core), nil
}
