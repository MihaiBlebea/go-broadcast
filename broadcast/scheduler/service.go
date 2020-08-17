package scheduler

// Task contains the spec and the func to be run
type Task struct {
	Spec string
	Func func()
}

// Service interface
type Service interface {
	Run(tasks []Task)
	Stop()
}
