package user

type UserRole string

var (
	UserRole_Admin UserRole = "ADMIN"
	UserRole_User  UserRole = "USER"
)

type User struct {
	ID       string
	Username string
	Password string
	Role     UserRole
}
