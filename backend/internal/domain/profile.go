package domain

type Profile struct {
	Username   string `json:"username"`
	FirstName  string `json:"firstName"`
	MiddleName string `json:"middleName"`
	LastName   string `json:"lastName"`
	Email      string `json:"email"`
	Role       int    `json:"role"`
}
