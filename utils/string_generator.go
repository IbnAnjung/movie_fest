package utils

import (
	enUtil "github.com/IbnAnjung/movie_fest/entity/utils"
	"github.com/google/uuid"
)

type stringGenerator struct {
	uuid uuid.UUID
}

func NewStringGenerator(uuid uuid.UUID) enUtil.StringGenerator {
	return stringGenerator{
		uuid: uuid,
	}
}

func (g stringGenerator) UUID() string {
	return g.uuid.String()
}
