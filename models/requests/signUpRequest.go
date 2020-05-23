package requests

type SignUp struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Secret   string `json:"secret"`
}
