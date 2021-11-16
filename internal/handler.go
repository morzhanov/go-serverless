package internal

import "log"

type Event struct {
	Name string `json:"name"`
}

type Handler interface {
	Handle(data []byte)
}

type handler struct {
}

func (s *handler) Handle(data []byte) {
	log.Printf("New Event handled\n")
	log.Printf("Data: %s\n", string(data))
}

func NewService() Handler {
	return &handler{}
}
