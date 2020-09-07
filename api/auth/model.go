package auth

type LoginRequest struct {
	User string `json:"user"`
	Pass string `json:"pass"`
}

type LoginResponse struct {
	Token string `json:"token"`
	Error string `json:"error,omitempty"`
}
