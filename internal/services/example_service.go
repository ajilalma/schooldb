package services

type ExampleService struct {
	repository ExampleRepository
}

func NewExampleService(repo ExampleRepository) *ExampleService {
	return &ExampleService{repository: repo}
}

func (s *ExampleService) PerformOperation() error {
	// Business logic goes here
	return nil
}