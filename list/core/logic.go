package core

import (
	"fmt"

	"github.com/MihaiBlebea/list/lists"
	"github.com/MihaiBlebea/list/sender"
	"github.com/sirupsen/logrus"
)

type service struct {
	listService   lists.Service
	senderService sender.Service
	logger        *logrus.Logger
}

// New returns a new core service
func New(
	listService lists.Service,
	senderService sender.Service,
	logger *logrus.Logger) Service {

	return &service{listService, senderService, logger}
}

func (s *service) Execute() error {
	contacts, err := s.listService.GetContacts()
	if err != nil {
		return err
	}

	for _, con := range *contacts {
		fmt.Println(con.Email)
		in := s.senderService.Build(con.Email, "<h1>This is an email</h1>", "Hello there!")
		out, err := s.senderService.Send(in)
		if err != nil {
			return err
		}

		s.logger.Info(out)
	}

	return nil
}
