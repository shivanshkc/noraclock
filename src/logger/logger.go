package logger

import (
	"errors"
	"fmt"
	"noraclock/v2/src/configs"
	"os"
	"path"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

var levelMap = map[string]zapcore.Level{
	"debug": zap.DebugLevel,
	"info":  zap.InfoLevel,
	"warn":  zap.WarnLevel,
	"error": zap.ErrorLevel,
}

// Get : Returns the uber-zap logger instance.
func Get() (*zap.Logger, error) {
	if logger != nil {
		return logger, nil
	}

	if err := initiateLogger(); err != nil {
		return nil, err
	}
	return logger, nil
}

func initiateLogger() error {
	fileEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	file, err := getLogFilePointer()
	if err != nil {
		return err
	}

	logLevel, exists := levelMap[configs.Logger.Level]
	if !exists {
		return errors.New("unknown log level")
	}

	levelEnabler := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= logLevel
	})

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, zapcore.Lock(file), levelEnabler),
		zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), levelEnabler),
	)

	logger = zap.New(core)
	return nil
}

func getLogFilePointer() (*os.File, error) {
	info, err := os.Stat(configs.Logger.File)
	if err == nil {
		if info.IsDir() {
			return nil, errors.New("provided path is a directory")
		}
		fmt.Println("Log file already present.")
		return os.OpenFile(configs.Logger.File, os.O_WRONLY, os.ModeAppend)
	}

	if !os.IsNotExist(err) {
		return nil, err
	}

	fmt.Println("Log file absent. Creating...")
	err = os.MkdirAll(path.Dir(configs.Logger.File), os.ModePerm)
	if err != nil && !os.IsExist(err) {
		return nil, errors.New("error while creating log file " + err.Error())
	}

	return os.Create(configs.Logger.File)
}
