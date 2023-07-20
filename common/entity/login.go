package entity

type Login struct {
	Username string
	Email    string
	Password string
}

func NewLogin(username, email, password string) *Login {
	return &Login{
		Username: username,
		Email:    email,
		Password: password,
	}
}
