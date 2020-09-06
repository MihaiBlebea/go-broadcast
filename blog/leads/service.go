package leads

// Service interface for leads service
type Service interface {
	Save(name, email string) error
}
