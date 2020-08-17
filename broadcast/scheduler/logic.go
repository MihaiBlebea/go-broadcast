package scheduler

import (
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

type service struct {
	cron *cron.Cron
}

// New returns a new scheduler service
func New(logger *logrus.Logger) Service {
	l := newLoggerAdaptor(logger)

	return &service{cron.New(cron.WithSeconds(), cron.WithLogger(l))}
}

func (s *service) Run(tasks []Task) {

	for _, task := range tasks {
		s.cron.AddFunc(task.Spec, task.Func)
	}

	s.cron.Start()
}

func (s *service) Stop() {
	s.cron.Stop()
}
