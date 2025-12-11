package response

type UserResponse struct {
	UUID  string `json:"uuid"`
	Email string `json:"email"`
	Token string `json:"token,omitempty"`
}