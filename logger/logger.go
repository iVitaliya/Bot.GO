package logger

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/andersfylling/disgord"
)

func InjectableLogger(conf zap.Config) *LoggerZap {
	log, err := conf.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}

	return &LoggerZap{
		instance: log.With(
			zap.String("lib", disgord.LibraryInfo()),
		),
	}
}

type LoggerZap struct {
	instance *zap.Logger
}

var _ disgord.Logger = (*LoggerZap)(nil)

func (log *LoggerZap) Debug(v ...interface{}) {
	log.instance.Debug(fmt.Sprint(v...))
	_ = log.instance.Sync()
}

func (log *LoggerZap) Info(v ...interface{}) {
	log.instance.Info(fmt.Sprint(v...))
	_ = log.instance.Sync()
}

func (log *LoggerZap) Error(v ...interface{}) {
	log.instance.Error(fmt.Sprint(v...))
	_ = log.instance.Sync()
}
