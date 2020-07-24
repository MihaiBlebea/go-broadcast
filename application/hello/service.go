package hello

// Service interface
type Service interface {
	BroadcastMessage(name string, age int, template string) (string, error)
}
