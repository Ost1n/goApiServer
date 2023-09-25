package entities

type EntityError struct{}

func (e *EntityError) Error() string {
	return "EntityError"
}
