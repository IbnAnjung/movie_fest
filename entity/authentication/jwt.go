package authentication

type UserJwtClaims struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}
