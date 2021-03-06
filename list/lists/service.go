package lists

// Service is a lists service interface
type Service interface {
	GetContacts() (*[]Contact, error)
	AddContact(name, email string) error
}
