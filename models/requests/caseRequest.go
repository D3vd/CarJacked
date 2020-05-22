package requests

// CaseRequest : Request body to create a new Case
type CaseRequest struct {
	User User `json:"user"`
	Car  Car  `json:"car"`
}

// User : User Information
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

// Car : Car Information
type Car struct {
	Color string `json:"color"`
	RegNo string `json:"regNo"`
}
