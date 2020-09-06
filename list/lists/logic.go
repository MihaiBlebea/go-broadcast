package lists

import (
	"net/http"

	"google.golang.org/api/people/v1"
)

type service struct {
	gService *people.Service
}

// Contact model
type Contact struct {
	Name  string
	Email string
}

// New returns a new lists service
func New(client *http.Client) (Service, error) {
	googleService, err := people.New(client)
	if err != nil {
		return &service{}, err
	}

	return &service{googleService}, nil
}

func (s *service) GetContacts() (*[]Contact, error) {
	var contacts []Contact

	resp, err := s.gService.
		People.
		Connections.
		List("people/me").
		PageSize(10).
		PersonFields("names,emailAddresses").
		Do()
	if err != nil {
		return &contacts, err
	}

	for _, contact := range resp.Connections {
		contacts = append(contacts, Contact{
			Name:  contact.Names[0].DisplayName,
			Email: contact.EmailAddresses[0].Value,
		})
	}

	return &contacts, nil
}

func (s *service) AddContact(name, email string) error {
	p := people.Person{}
	p.Names = append(p.Names, &people.Name{
		DisplayName: name,
	})

	p.EmailAddresses = append(p.EmailAddresses, &people.EmailAddress{
		Value: email,
	})

	call := s.gService.People.CreateContact(&p)
	_, err := call.Do()
	if err != nil {
		return err
	}

	return nil
}
