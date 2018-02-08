package mail

import "fmt"

// SendError allows access to the underlying error given during a faulty send
type SendError struct {
	Index uint
	Cause error
}

func (err *SendError) Error() string {
	return fmt.Sprintf("gomail: could not send email %d: %v",
		err.Index, err.Cause)
}
