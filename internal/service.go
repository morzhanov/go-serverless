package internal

type Service interface {
	Handle() error
}

type service struct {
}

func (s *service) Handle() error {
	return nil
}

func NewService() Service {
	return &service{}
}
