package domain

type User struct {
	username string
	password string
	coins    int
}

func NewUser(username, password string) *User {
	return &User{
		username: username,
		password: password,
	}
}

func (user User) GetUsername() string {
	return user.username
}

func (user User) GetPassword() string {
	return user.password
}
