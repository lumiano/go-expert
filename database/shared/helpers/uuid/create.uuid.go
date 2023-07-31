package uuid

import "github.com/google/uuid"

type Helpers interface {
	Create() string
}

type UUID struct {
}

func (u UUID) Create() string {
	return uuid.New().String()
}
