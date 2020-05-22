package requests

// Login : Request body for Login
type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
