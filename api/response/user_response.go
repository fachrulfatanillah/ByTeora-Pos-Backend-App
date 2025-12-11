package response

type UserResponse struct {
	UUID  string `json:"uuid"`
	Email string `json:"email"`
	Role  string `json:"role,omitempty"`
	Token string `json:"token,omitempty"`
}