package core

import (
	"strconv"

	"github.com/rs/zerolog/log"
)

func MustAtoi(input string) int {
	intValue, err := strconv.Atoi(input)
	if err != nil {
		log.Panic().Err(err).Msg("Error converting string to int")
	}

	return intValue
}
