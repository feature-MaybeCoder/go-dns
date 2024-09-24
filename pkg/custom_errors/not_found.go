package custom_errors

import "fmt"

type NotFoundError struct {
	ObjName string
}

func (err NotFoundError) Error() string {
	return fmt.Sprintf("No such %s", err.ObjName)
}
