package authentication

type Register struct {
	Username        string
	Password        string
	ConfirmPassword string
}

type RegisteredUser struct {
	ID       int64
	Username string
	Token    string
}
