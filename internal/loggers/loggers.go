package loggers

import (
	"github.com/ddreval/waytogo/internal/config"
	"github.com/samber/do"
	"github.com/sirupsen/logrus"
)

func New(di *do.Injector) (*logrus.Logger, error) {
	logger := logrus.StandardLogger()
	logger.SetReportCaller(true)
	cfg, err := do.Invoke[*config.Config](di)
	if err != nil {
		return nil, err
	}
	for _, level := range logrus.AllLevels {
		if level.String() == cfg.LogLevel {
			logger.SetLevel(level)
			break
		}
	}
	return logger, nil
}
