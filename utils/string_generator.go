package utils

import (
	enUtil "github.com/IbnAnjung/movie_fest/entity/utils"
	"github.com/google/uuid"
)

type stringGenerator struct {
}

func NewStringGenerator() enUtil.StringGenerator {
	return stringGenerator{}
}

func (g stringGenerator) UUID() string {
	return uuid.New().String()
}
