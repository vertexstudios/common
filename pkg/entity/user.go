package entity

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}

func NewUser(username, email, name string) *User {
	return &User{
		Username: username,
		Email:    email,
		Name:     name,
	}
}
