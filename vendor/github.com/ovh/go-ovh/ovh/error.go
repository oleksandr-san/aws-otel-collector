package ovh

<<<<<<< HEAD
import "fmt"
=======
import (
	"fmt"
	"strings"
)
>>>>>>> main

// APIError represents an error that can occurred while calling the API.
type APIError struct {
	// Error class
	Class string `json:"class,omitempty"`
	// Error message.
	Message string `json:"message"`
	// Error details
	Details map[string]string `json:"details,omitempty"`
	// HTTP code.
	Code int
	// ID of the request
	QueryID string
}

<<<<<<< HEAD
func (err *APIError) Error() string {
	if err.Class == "" {
		return fmt.Sprintf("HTTP Error %d: %q", err.Code, err.Message)
	}

	return fmt.Sprintf("HTTP Error %d: %s: %q (X-OVH-Query-Id: %s)", err.Code, err.Class, err.Message, err.QueryID)
=======
// Let's make sure that APIError always satisfies the fmt.Stringer and error interfaces
var _ fmt.Stringer = APIError{}
var _ error = APIError{}

func (err APIError) Error() string {
	var sb strings.Builder
	sb.Grow(128)

	// Base message
	fmt.Fprint(&sb, "OVHcloud API error (status code ", err.Code, "): ")

	// Append class if any
	if err.Class != "" {
		fmt.Fprint(&sb, err.Class, ": ")
	}

	// Real error message, quoted
	fmt.Fprintf(&sb, "%q", err.Message)

	// QueryID if any
	if err.QueryID != "" {
		fmt.Fprint(&sb, " (X-OVH-Query-Id: ", err.QueryID, ")")
	}

	return sb.String()
}

func (err APIError) String() string {
	return err.Error()
>>>>>>> main
}
