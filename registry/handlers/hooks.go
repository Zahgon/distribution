package handlers

import (
	"github.com/sirupsen/logrus"
)

type logHook struct {
	LevelsParam []string
	Mail        *mailer
}

func (hook *logHook) Fire(entry *logrus.Entry) error { _ = "STUB: not implemented"; return nil }

func (hook *logHook) Levels() []logrus.Level { _ = "STUB: not implemented"; return nil }
