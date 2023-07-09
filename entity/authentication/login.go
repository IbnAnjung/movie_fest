package authentication

type Login struct {
	Username string
	Password string
}

type LogedinUser struct {
	ID       int64
	Username string
	Token    string
}
