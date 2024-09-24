package custom_errors

import "fmt"

type UnexpectedBehaviourError struct {
	Message string
}

func (err UnexpectedBehaviourError) Error() string {
	return fmt.Sprintf("Unexpected behaviour: %s", err.Message)
}
