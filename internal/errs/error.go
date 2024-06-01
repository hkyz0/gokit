package errs

import "fmt"

func NewErrIndexOutOfRang(index int) error {
	return fmt.Errorf("index %d out of range", index)
}

func NewErrInvalidType(want string) error {
	return fmt.Errorf("type %s not recognized", want)
}
