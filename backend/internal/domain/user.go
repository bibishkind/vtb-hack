package domain

const (
	UserRoleMember = iota
	UserRoleLeader
	UserRoleEditor
	UserRoleAdmin
)

type User struct {
	Id         int    `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	FirstName  string `json:"firstName"`
	MiddleName string `json:"middleName"`
	LastName   string `json:"lastName"`
	Email      string `json:"email"`
	Role       int    `json:"role"`
	Wallet     string `json:"wallet"`
}
