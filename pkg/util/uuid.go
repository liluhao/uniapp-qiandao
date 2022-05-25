package util

import "github.com/google/uuid"

func GetUUID() string {
	random, _ := uuid.NewRandom()
	return random.String()
}
