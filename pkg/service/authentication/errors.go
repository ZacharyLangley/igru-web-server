package authentication

import "fmt"

type InvalidUserIDError struct {
	UserID string
}

func (i InvalidUserIDError) Error() string {
	return fmt.Sprintf("invalid user ID specified: %s", i.UserID)
}
