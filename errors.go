package mail

import "fmt"

// SendError represents the failure to transmit a Message contains the cause of failure
// determined by the internal error.
type SendError struct {
	Index uint //Index represents the index of the message which could not be transmitted
	Cause error
}

func (err *SendError) Error() string {

	return fmt.Sprintf("gomail: could not send email %d: %v",
		err.Index, err.Cause)
}
