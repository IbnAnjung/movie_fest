package authentication

import "fmt"

type UserJwtClaims struct {
	ID       string `json:"id"`
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
}

func (c *UserJwtClaims) GenerateID(randomString string) {
	c.ID = fmt.Sprintf("user_token:%d-%s", c.UserID, randomString)
}

func (c *UserJwtClaims) GetID() string {
	return c.ID
}

type JwtKey string

const JwtKey_User JwtKey = "User"
