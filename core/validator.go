package core

type Validatable interface {
	IsValid() error
}
