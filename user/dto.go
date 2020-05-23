package user

// User - user data transfer object
type User struct {
	ID        int    `json:"id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Phone     string `json:"phone,omitempty"`
}

// ErrorMessage - error message when an error happens
type ErrorMessage struct {
	Message string `json:"error,omitempty"`
}
