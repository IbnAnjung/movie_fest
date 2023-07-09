package user

type UserRole string

var (
	UserRole_Admin UserRole = "ADMIN"
	UserRole_User  UserRole = "USER"
)

type User struct {
	ID       int64
	Username string
	Password string
}

type UserToken struct {
	ID      int64
	UserID  int64
	Token   string
	IsBlock bool
}
